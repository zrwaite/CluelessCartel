package models

import "go.mongodb.org/mongo-driver/bson"

type Weapons struct {
	Guns       int
	Explosives int
	Ammunition int
}

var weaponsStruct = bson.M{
	"bsonType": "object",
	"required": []string{"guns", "explosives", "ammunition"},
	"properties": bson.M{
		"guns":       bson.M{"bsonType": "int"},
		"explosives": bson.M{"bsonType": "int"},
		"ammunition": bson.M{"bsonType": "int"},
	},
}
