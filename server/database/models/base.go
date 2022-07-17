package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Base struct {
	Location    Location
	Index       int
	HexagonRows []HexagonRow `bson:"hexagon_rows"`
}

var baseSchema = bson.M{
	"bsonType": "object",
	"required": []string{"location", "index", "hexagon_rows"},
	"properties": bson.M{
		"location":     locationSchema,
		"index":        bsonInt,
		"hexagon_rows": bsonArray(hexagonRowsSchema),
	},
}
