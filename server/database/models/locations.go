package models

import "go.mongodb.org/mongo-driver/bson"

type Location struct {
	Name         string
	LandMaterial LandMaterial `bson:"landMaterial"`
}

var locationSchema = bson.M{
	"bsonType": "object",
	"required": []string{"name", "landMaterial"},
	"properties": bson.M{
		"name":         bsonString,
		"landMaterial": landMaterialSchema,
	},
}
