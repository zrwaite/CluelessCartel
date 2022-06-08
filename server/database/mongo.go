package database

import (
	"clueless-cartel-server/settings"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoDatabase *mongo.Database

func ConnectToMongoDB() {
	var mongoUrl string
	if settings.DEV {
		mongoUrl = os.Getenv("DEV_MONGO_URL")
	} else {
		mongoUrl = os.Getenv("PROD_MONGO_URL")
	}
	clientOptions := options.Client().ApplyURI(mongoUrl)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	if settings.DEV {
		MongoDatabase = client.Database("CluelessCartelDevelopmentCluster")
	} else {
		MongoDatabase = client.Database("CluelessCartelProductionCluster")
	}
}
