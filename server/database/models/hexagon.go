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
	Owned        bool
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
	var hexagonRows = []HexagonRow{}
	for i := -2; i <= 2; i++ {
		var hexagonRow = HexagonRow{Index: i}
		var startIndex int
		var endIndex int
		if i == 2 || i == -2 {
			startIndex = -1
			endIndex = 1
		} else if i == 1 || i == -1 {
			startIndex = -2
			endIndex = 1
		} else {
			startIndex = -2
			endIndex = 2
		}
		for i2 := startIndex; i2 <= endIndex; i2++ {
			owned := true
			if i == -2 || i == 2 || i2 == startIndex || i2 == endIndex {
				owned = false
			}
			hexagonRow.Hexagons = append(hexagonRow.Hexagons, Hexagon{
				LandMaterial: landMaterial,
				Index:        i2,
				Structure:    EmptyStructure,
				Owned:        owned,
			})
		}
		hexagonRows = append(hexagonRows, hexagonRow)
	}
	return hexagonRows
}
