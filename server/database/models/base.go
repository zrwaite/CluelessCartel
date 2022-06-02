package models

import "go.mongodb.org/mongo-driver/bson"

var Location = struct {
	Nevada         string
	Florida        string
	CanadianBorder string
	NewYork        string
}{"Nevada", "Florida", "Canadian Border", "New York"}

type Base struct {
	Location    string
	Resources   Resources
	Drugs       Drugs
	Weapons     Weapons
	HexagonRows []HexagonRow `bson:"hexagon_rows"`
}

func CreateStartingBase(location string) Base {
	var material string
	switch location {
	case "Nevada":
		material = "Sand"
	case "Florida":
		material = "Sand"
	case "Canadian Border":
		material = "Grass"
	case "New York":
		material = "Pavement"
	default:
		material = "Dirt"
	}
	return Base{
		Location:    location,
		Resources:   StartingResources,
		Drugs:       StartingDrugs,
		Weapons:     StartingWeapons,
		HexagonRows: CreateStartingHexagonRows(material),
	}

}

var baseSchema = bson.M{
	"bsonType": "object",
	"required": []string{"resources", "drugs", "weapons", "hexagon_rows"},
	"properties": bson.M{
		"resources": resourcesSchema,
		"drugs":     drugsSchema,
		"weapons":   weaponsStruct,
		"hexagon_rows": bson.M{
			"bsonType":    "array",
			"uniqueItems": false,
			"items":       hexagonRowsSchema,
		},
	},
}
