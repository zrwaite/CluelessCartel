package user

import (
	"clueless-cartel-server/api"
	"clueless-cartel-server/database"
	"clueless-cartel-server/database/models"
	"context"

	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getUser(r *http.Request, res *api.Response) {
	username := r.URL.Query().Get("username")
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
	opts := options.FindOne().SetProjection(models.GetUserQuery)
	cursor := usersCollection.FindOne(context.TODO(), filter, opts)
	if err := cursor.Decode(&models.GetUserReturn); err != nil {
		res.Errors = append(res.Errors, "Failed to get database data - "+err.Error())
	} else {
		res.Success = true
		res.Status = 200
		res.Response = models.GetUserReturn
	}
}
