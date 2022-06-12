package models

import (
	"math/rand"

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

var Grass = LandMaterial{"Grass"}
var Dirt = LandMaterial{"Dirt"}
var Sand = LandMaterial{"Sand"}
var Pavement = LandMaterial{"Pavement"}
var Water = LandMaterial{"Water"}
var AllLandMaterials = []LandMaterial{Grass, Dirt, Sand, Pavement, Water}

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

func ValidateHexagonRowsStructure(hexagonRows *[]HexagonRow) (valid bool) {
	if len(*hexagonRows) < 5 {
		return false
	}
	length := len((*hexagonRows)[0].Hexagons)
	if length < 5 {
		return false
	}
	for _, hexagonRow := range *hexagonRows {
		newLength := len(hexagonRow.Hexagons)
		if length != newLength {
			return false
		}
	}
	return true
}

func CreateStartingHexagonRows(location Location) []HexagonRow {
	var hexagonRows = []HexagonRow{}
	for i1 := -2; i1 <= 2; i1++ {
		/* Create hexagon Grid:
		  ❌ ⬜ ⬜ ⬜ ❌
		   ⬜ ⬛ ⬛ ⬜ ❌
		 ⬜ ⬛ ⬛ ⬛ ⬜
		  ⬜ ⬛ ⬛ ⬜ ❌
		❌ ⬜ ⬜ ⬜ ❌
		*/
		var hexagonRow = HexagonRow{Y: i1}
		for i2 := -2; i2 <= 2; i2++ {
			owned := false
			buyable := true
			if i1 == -2 || i1 == 2 {
				if i2 == -2 || i2 == 2 {
					buyable = false
				}
			} else if i1 == -1 || i1 == 1 {
				if i2 == 2 {
					buyable = false
				} else if i2 == -1 || i2 == 0 {
					owned = true
				}
			} else if i2 == -1 || i2 == 0 || i2 == 1 {
				owned = true
			}
			hexagonRow.Hexagons = append(hexagonRow.Hexagons, Hexagon{
				LandMaterial: location.LandMaterial,
				Structure:    EmptyStructure,
				Owned:        owned,
				Rotation:     rand.Intn(6),
				Buyable:      buyable,
				X:            i2,
			})
		}
		hexagonRows = append(hexagonRows, hexagonRow)
	}
	return hexagonRows
}
