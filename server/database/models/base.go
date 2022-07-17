package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Base struct {
	Location    Location
	Index       int
	Resources   []ResourcesAmount
	Drugs       []ResourcesAmount
	Weapons     Weapons
	HexagonRows []HexagonRow `bson:"hexagon_rows"`
}

var baseSchema = bson.M{
	"bsonType": "object",
	"required": []string{"location", "index", "resources", "drugs", "weapons", "hexagon_rows"},
	"properties": bson.M{
		"location":     locationSchema,
		"index":        bsonInt,
		"resources":    bsonArray(resourcesAmountSchema),
		"drugs":        bsonArray(resourcesAmountSchema),
		"weapons":      weaponsStruct,
		"hexagon_rows": bsonArray(hexagonRowsSchema),
	},
}
