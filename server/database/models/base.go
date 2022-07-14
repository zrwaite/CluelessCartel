package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Base struct {
	Location    Location
	Index       int
	Resources   []ResourcesAmount
	Drugs       []DrugAmount
	Weapons     Weapons
	HexagonRows []HexagonRow `bson:"hexagon_rows"`
}

var baseSchema = bson.M{
	"bsonType": "object",
	"required": []string{"resources", "drugs", "weapons", "hexagon_rows"},
	"properties": bson.M{
		"resources": bson.M{
			"bsonType":    "array",
			"uniqueItems": false,
			"items":       resourcesAmountSchema,
		},
		"drugs": bson.M{
			"bsonType":    "array",
			"uniqueItems": false,
			"items":       drugAmountSchema,
		},
		"weapons": weaponsStruct,
		"hexagon_rows": bson.M{
			"bsonType":    "array",
			"uniqueItems": false,
			"items":       hexagonRowsSchema,
		},
	},
}
