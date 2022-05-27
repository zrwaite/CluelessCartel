package user

import (
	"clueless-cartel-server/api"
	"clueless-cartel-server/database"
	"context"

	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getUser(r *http.Request, res *api.Response) {
	username := r.URL.Query().Get("username")
	var getProjection = bson.D{{"username", 1}}

	if username == "" {
		res.Errors = append(res.Errors, "Username not defined")
		return
	}
	filter := bson.D{{
		Key: "username",
		Value: bson.D{{
			Key:   "$eq",
			Value: username,
		}},
	}}
	usersCollection := database.MongoDatabase.Collection("users")
	opts := options.FindOne().SetProjection(getProjection)
	cursor := usersCollection.FindOne(context.TODO(), filter, opts)
	var result bson.M
	if err := cursor.Decode(&result); err != nil {
		res.Errors = append(res.Errors, "Failed to get database data - "+err.Error())
	} else {
		res.Success = true
		res.Status = 200
		res.Response = result
	}
}
