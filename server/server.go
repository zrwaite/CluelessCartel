package main

import (
	"clueless-cartel-server/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client

type Response struct {
	Success  bool
	Errors   []string
	Message  string
	Response interface{}
}

type Hello struct {
	Name string
	Age  int
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		helloGetHandler(w, r)
	} else if r.Method == "POST" {
		helloPostHandler(w, r)
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}
func helloGetHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{true, []string{}, "hello", nil}
	filter := bson.D{{
		Key: "age",
		Value: bson.D{{
			Key:   "$gt",
			Value: 25,
		}},
	}}
	usersCollection := mongoClient.Database("CluelessCartelProductionCluster").Collection("users")
	cursor, err := usersCollection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
	json.NewEncoder(w).Encode(res)
}
func helloPostHandler(w http.ResponseWriter, r *http.Request) {
	var h Hello
	res := Response{false, []string{}, "", nil}
	err := json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		res.Errors = append(res.Errors, (err.Error()))
	} else {
		if h.Name == "" {
			res.Errors = append(res.Errors, "Missing name")
		} else {
			res.Success = true
		}
	}
	if res.Success {
		res.Message = "hello " + h.Name
		w.WriteHeader(http.StatusCreated)
	}
	json.NewEncoder(w).Encode(res)

	/*

		usersCollection := mongoClient.Database("CluelessCartelProductionCluster").Collection("users")
			users := []interface{}{
				bson.D{{"fullName", "User 2"}, {"age", 25}},
				bson.D{{"fullName", "User 3"}, {"age", 20}},
				bson.D{{"fullName", "User 4"}, {"age", 28}},
			}
			result, err := usersCollection.InsertMany(context.TODO(), users)
			if err != nil {
				panic(err)
			}

			fmt.Println(result.InsertedIDs)

	*/
}

func main() {
	mongoClient = database.ConnectToMongoDB()
	fileServer := http.FileServer(http.Dir("../client"))
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/", fileServer)
	fmt.Printf("Starting server at port 6969 with hello\n")
	if err := http.ListenAndServe(":6969", nil); err != nil {
		log.Fatal(err)
	}
}
