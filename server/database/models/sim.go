package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Sim struct {
	Name  string
	IQ    int
	EQ    int
	Alive bool
	XP    int
}

var simSchema = bson.M{
	"bsonType": "object",
	"required": []string{"name", "iq", "eq", "alive", "xp"},
	"properties": bson.M{
		"name":  bsonString,
		"iq":    bsonInt,
		"eq":    bsonInt,
		"alive": bsonBoolean,
		"xp":    bsonInt,
	},
}
