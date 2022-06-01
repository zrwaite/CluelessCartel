package models

import "go.mongodb.org/mongo-driver/bson"

var LandMaterial = struct {
	Dirt     int
	Grass    int
	Sand     int
	Pavement int
}{0, 1, 2, 3}

type Hexagon struct {
	LandMaterial int `bson:"land_material"`
	Structure    Structure
	Index        int
}

var hexagonsSchema = bson.M{
	"bsonType": "object",
	"required": []string{"land_material", "structure", "index"},
	"properties": bson.M{
		"land_material": bson.M{"bsonType": "int"},
		"structure":     structureSchema,
		"index":         bson.M{"bsonType": "int"},
	},
}

type HexagonRow struct {
	Index    int
	Hexagons []Hexagon
}

var hexagonRowsSchema = bson.M{
	"bsonType": "object",
	"required": []string{"index", "hexagons"},
	"properties": bson.M{
		"index": bson.M{"bsonType": "int"},
		"hexagons": bson.M{
			"bsonType":    "array",
			"uniqueItems": false,
			"items":       hexagonsSchema,
		},
		"resource_capacity": resourcesSchema,
	},
}

var StartingHexagonRows = [3]HexagonRow{
	{
		Index: -1,
		Hexagons: []Hexagon{
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        -1,
				Structure:    EmptyStructure,
			},
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        0,
				Structure:    EmptyStructure,
			},
		},
	},
	{
		Index: 0,
		Hexagons: []Hexagon{
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        -1,
				Structure:    EmptyStructure,
			},
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        0,
				Structure:    EmptyStructure,
			},
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        1,
				Structure:    EmptyStructure,
			},
		},
	},
	{
		Index: 1,
		Hexagons: []Hexagon{
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        -1,
				Structure:    EmptyStructure,
			},
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        0,
				Structure:    EmptyStructure,
			},
		},
	},
}
