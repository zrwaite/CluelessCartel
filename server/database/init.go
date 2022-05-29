package database

import (
	"clueless-cartel-server/database/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeDatabase() {
	MongoDatabase.Collection("users").Drop(context.TODO())
	var options = new(options.CreateCollectionOptions)
	options.Validator = models.UsersValidator
	error := MongoDatabase.CreateCollection(context.TODO(), "users", options)
	if error != nil {
		log.Fatal("Failed to create users collection" + error.Error())
	}
}
