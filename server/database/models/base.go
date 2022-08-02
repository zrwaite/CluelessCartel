package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Base struct {
	Location    Location
	Index       int
	HexagonRows []HexagonRow `bson:"hexagonRows"`
	Tasks       []Task
	Sims        []Sim
}

var baseSchema = bson.M{
	"bsonType": "object",
	"required": []string{"location", "index", "hexagonRows", "tasks", "sims"},
	"properties": bson.M{
		"location":    locationSchema,
		"index":       bsonInt,
		"hexagonRows": bsonArray(hexagonRowsSchema),
		"tasks":       bsonArray(taskSchema),
		"sims":        bsonArray(simSchema),
	},
}
