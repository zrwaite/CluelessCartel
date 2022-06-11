package models

import (
	"errors"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
)

type Base struct {
	Location    Location
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

func CreateStartingBase(locationName string, index int) (Base, error) {
	var base Base
	location, success := GetLocation(locationName)
	if !success {
		return base, errors.New("invalid location")
	}
	base = Base{
		Location:    location,
		Index:       index,
		Resources:   StartingResources,
		Drugs:       StartingDrugs,
		Weapons:     StartingWeapons,
		HexagonRows: CreateStartingHexagonRows(location),
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
