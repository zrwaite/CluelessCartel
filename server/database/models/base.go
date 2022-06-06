package models

import (
	"errors"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
)

var Location = struct {
	Nevada         string
	Florida        string
	CanadianBorder string
	NewYork        string
}{"Nevada", "Florida", "Canadian Border", "New York"}

type Base struct {
	Location    string
	Index       int
	Resources   Resources
	Drugs       Drugs
	Weapons     Weapons
	HexagonRows []HexagonRow `bson:"hexagon_rows"`
}

func GetRandomLandMaterial() (material string) {
	materials := []string{"Sand", "Grass", "Dirt", "Pavement"}
	return materials[rand.Intn(len(materials))]
}

func GetLandMaterial(location string) (material string) {
	switch location {
	case "Nevada":
		material = "Sand"
	case "Florida":
		material = "Sand"
	case "Canadian Border":
		material = "Grass"
	case "New York":
		material = "Pavement"
	}
	return
}

func CreateStartingBase(location string, index int) (Base, error) {
	var base Base
	material := GetLandMaterial(location)
	if material == "" {
		return base, errors.New("invalid location")
	}

	base = Base{
		Location:    location,
		Index:       index,
		Resources:   StartingResources,
		Drugs:       StartingDrugs,
		Weapons:     StartingWeapons,
		HexagonRows: CreateStartingHexagonRows(material),
	}
	return base, nil

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
