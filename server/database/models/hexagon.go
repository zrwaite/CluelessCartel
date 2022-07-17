package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type LandMaterial struct {
	Name string
}

var landMaterialSchema = bson.M{
	"bsonType": "object",
	"required": []string{"name"},
	"properties": bson.M{
		"name": bsonString,
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
		"land_material": landMaterialSchema,
		"structure":     structureSchema,
		"x":             bsonInt,
		"rotation":      bsonInt,
		"owned":         bsonBoolean,
		"buyable":       bsonBoolean,
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
		"y":        bsonInt,
		"hexagons": bsonArray(hexagonsSchema),
	},
}
