package models

import "go.mongodb.org/mongo-driver/bson"

type Drugs struct {
	Opioids int
	Weed    int
	Meth    int
}

var drugsSchema = bson.M{
	"bsonType": "object",
	"required": []string{"opioids", "weed", "meth"},
	"properties": bson.M{
		"opioids": bson.M{"bsonType": "int"},
		"weed":    bson.M{"bsonType": "int"},
		"meth":    bson.M{"bsonType": "int"},
	},
}

var StartingDrugs = Drugs{10, 10, 0}
