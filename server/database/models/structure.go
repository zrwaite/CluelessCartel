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
		"moveable":          bsonBoolean,
		"natural":           bsonBoolean,
		"enemy":             bsonBoolean,
		"name":              bsonString,
		"land_materials":    bsonArray(landMaterialSchema),
		"resource_capacity": bsonArray(resourcesAmountSchema),
	},
}
