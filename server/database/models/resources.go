package models

import "go.mongodb.org/mongo-driver/bson"

type Resource struct {
	Name string
}

type ResourcesAmount struct {
	ResourceName string
	Amount       int
}

type ResourceCost struct {
	ResourceName string
	Cost         int
}

var resourcesAmountSchema = bson.M{
	"bsonType": "object",
	"required": []string{"resource_name", "amount"},
	"properties": bson.M{
		"resource_name": bson.M{"bsonType": "string"},
		"amount":        bson.M{"bsonType": "int"},
	},
}
