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
	X            int
	Owned        bool
	Buyable      bool
}

var hexagonsSchema = bson.M{
	"bsonType": "object",
	"required": []string{"land_material", "structure", "owned", "buyable"},
	"properties": bson.M{
		"land_material": bson.M{"bsonType": "string"},
		"structure":     structureSchema,
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
	"required": []string{"starting_index", "hexagons"},
	"properties": bson.M{
		"hexagons": bson.M{
			"bsonType":    "array",
			"uniqueItems": false,
			"items":       hexagonsSchema,
		},
	},
}

func ValidateHexagonRowsStructure(hexagonRows []HexagonRow) (valid bool) {
	if len(hexagonRows) < 5 {
		return false
	}
	length := len(hexagonRows[0].Hexagons)
	if length < 5 {
		return false
	}
	for _, hexagonRow := range hexagonRows {
		newLength := len(hexagonRow.Hexagons)
		if length != newLength {
			return false
		}
	}
	return true
}

func CreateStartingHexagonRows(landMaterial string) []HexagonRow {
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
				LandMaterial: landMaterial,
				Structure:    EmptyStructure,
				Owned:        owned,
				Buyable:      buyable,
				X:            i2,
			})
		}
		hexagonRows = append(hexagonRows, hexagonRow)
	}
	return hexagonRows
}
