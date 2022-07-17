package models

import "go.mongodb.org/mongo-driver/bson"

type Location struct {
	Name         string
	LandMaterial LandMaterial `bson:"land_material"`
}

var locationSchema = bson.M{
	"bsonType": "object",
	"required": []string{"name", "land_material"},
	"properties": bson.M{
		"name":          bsonString,
		"land_material": landMaterialSchema,
	},
}
