package models

import "go.mongodb.org/mongo-driver/bson"

type Weapons struct {
	Guns       int
	Explosives int
}

var weaponsStruct = bson.M{
	"bsonType": "object",
	"required": []string{"guns", "explosives"},
	"properties": bson.M{
		"guns":       bson.M{"bsonType": "int"},
		"explosives": bson.M{"bsonType": "int"},
	},
}

var StartingWeapons = Weapons{1, 0}
