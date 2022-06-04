package models

import "go.mongodb.org/mongo-driver/bson"

type Drug struct {
	Cost        []ResourceCost
	BatchAmount int
	UnitPrice   int
}

var Opioids = Drug{
	Cost:        []ResourceCost{{Plants, 5}, {Chemicals, 5}},
	BatchAmount: 20,
	UnitPrice:   50,
}

var Weed = Drug{
	Cost:        []ResourceCost{{Plants, 10}},
	BatchAmount: 50,
	UnitPrice:   16,
}

var Meth = Drug{
	Cost:        []ResourceCost{{Chemicals, 10}},
	BatchAmount: 40,
	UnitPrice:   20,
}

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
