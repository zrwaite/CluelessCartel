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
	ResourceName string `bson:"resource_name"`
	Amount       int
	Quality		 float64 //percentage
}

var resourcesAmountSchema = bson.M{
	"bsonType": "object",
	"required": []string{"resource_name", "amount", "quality"},
	"properties": bson.M{
		"resource_name": bsonString,
		"amount":        bsonInt,
		"quality":       bsonDouble,
	},
}
