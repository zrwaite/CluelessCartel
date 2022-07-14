package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Structure struct {
	Moveable         bool
	Enemy            bool
	Natural          bool
	Name             string
	LandMaterials    []LandMaterial    `bson:"land_materials"`
	ResourceCapacity []ResourcesAmount `bson:"resource_capacity"`
}

var structureSchema = bson.M{
	"bsonType": "object",
	"required": []string{"moveable", "enemy", "natural", "name", "land_materials", "resource_capacity"},
	"properties": bson.M{
		"moveable": bson.M{"bsonType": "bool"},
		"natural":  bson.M{"bsonType": "bool"},
		"enemy":    bson.M{"bsonType": "bool"},
		"name":     bson.M{"bsonType": "string"},
		"land_materials": bson.M{
			"bsonType":    "array",
			"uniqueItems": true,
			"items":       landMaterialScema,
		},
		"resource_capacity": bson.M{
			"bsonType":    "array",
			"uniqueItems": false,
			"items":       resourcesAmountSchema,
		},
	},
}
