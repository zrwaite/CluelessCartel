package models

import "go.mongodb.org/mongo-driver/bson"

var LandMaterial = struct {
	Dirt     string
	Grass    string
	Sand     string
	Pavement string
}{"Dirt", "Grass", "Sand", "Pavement"}

type Hexagon struct {
	LandMaterial string `bson:"land_material"`
	Structure    Structure
	Index        int
}

var hexagonsSchema = bson.M{
	"bsonType": "object",
	"required": []string{"land_material", "structure", "index"},
	"properties": bson.M{
		"land_material": bson.M{"bsonType": "string"},
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

func CreateStartingHexagonRows(landMaterial string) []HexagonRow {
	return []HexagonRow{
		{
			Index: -1,
			Hexagons: []Hexagon{
				{
					LandMaterial: landMaterial,
					Index:        -1,
					Structure:    EmptyStructure,
				},
				{
					LandMaterial: landMaterial,
					Index:        0,
					Structure:    EmptyStructure,
				},
			},
		},
		{
			Index: 0,
			Hexagons: []Hexagon{
				{
					LandMaterial: landMaterial,
					Index:        -1,
					Structure:    EmptyStructure,
				},
				{
					LandMaterial: landMaterial,
					Index:        0,
					Structure:    EmptyStructure,
				},
				{
					LandMaterial: landMaterial,
					Index:        1,
					Structure:    EmptyStructure,
				},
			},
		},
		{
			Index: 1,
			Hexagons: []Hexagon{
				{
					LandMaterial: landMaterial,
					Index:        -1,
					Structure:    EmptyStructure,
				},
				{
					LandMaterial: landMaterial,
					Index:        0,
					Structure:    EmptyStructure,
				},
			},
		},
	}
}
