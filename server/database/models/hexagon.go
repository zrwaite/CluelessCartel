package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type LandMaterial struct {
	Name string
}

var landMaterialScema = bson.M{
	"bsonType": "object",
	"required": []string{"name"},
	"properties": bson.M{
		"name": bson.M{"bsonType": "string"},
	},
}

type Hexagon struct {
	LandMaterial LandMaterial `bson:"land_material"`
	Structure    Structure
	Rotation     int
	X            int
	Owned        bool
	Buyable      bool
}

var hexagonsSchema = bson.M{
	"bsonType": "object",
	"required": []string{"land_material", "rotation", "structure", "x", "owned", "buyable"},
	"properties": bson.M{
		"land_material": landMaterialScema,
		"structure":     structureSchema,
		"x":             bson.M{"bsonType": "int"},
		"rotation":      bson.M{"bsonType": "int"},
		"owned":         bson.M{"bsonType": "bool"},
		"buyable":       bson.M{"bsonType": "bool"},
	},
}

type HexagonRow struct {
	Hexagons []Hexagon
	Y        int
}

var hexagonRowsSchema = bson.M{
	"bsonType": "object",
	"required": []string{"y", "hexagons"},
	"properties": bson.M{
		"y": bson.M{"bsonType": "int"},
		"hexagons": bson.M{
			"bsonType":    "array",
			"uniqueItems": false,
			"items":       hexagonsSchema,
		},
	},
}
