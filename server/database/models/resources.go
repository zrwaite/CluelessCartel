package models

import "go.mongodb.org/mongo-driver/bson"

type Resource struct {
	Name               string
	SynthesizationCost []ResourcesAmount
	SynthesizationTime int // in minutes
	BatchAmount        int
	StartingUnitPrice  int
}

type ResourcesAmount struct {
	ResourceName string
	Amount       int
}

var resourcesAmountSchema = bson.M{
	"bsonType": "object",
	"required": []string{"name", "amount"},
	"properties": bson.M{
		"name":   bson.M{"bsonType": "string"},
		"amount": bson.M{"bsonType": "int"},
	},
}
