package hello

import (
	"clueless-cartel-server/api"
	"clueless-cartel-server/database"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type Hello struct {
	Name string
	Age  int
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	api.AddJSONHeader(w)
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
	res := new(api.Response).Init()
	filter := bson.D{{
		Key: "age",
		Value: bson.D{{
			Key:   "$gt",
			Value: 25,
		}},
	}}
	usersCollection := database.MongoDatabase.Collection("users")
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
	res := new(api.Response).Init()
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

	usersCollection := database.MongoDatabase.Collection("users")
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
}
