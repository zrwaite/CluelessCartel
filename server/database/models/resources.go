package models

import "go.mongodb.org/mongo-driver/bson"

type ResourceCost struct {
	Resource Resource
	Cost     int
}

type Resource struct {
	Name string
}

var Metal = Resource{
	Name: "Metal",
}

var Plants = Resource{
	Name: "Plants",
}

var Chemicals = Resource{
	Name: "Chemicals",
}

type Resources struct {
	Metal     int
	Plants    int
	Chemicals int
}

var resourcesSchema = bson.M{
	"bsonType": "object",
	"required": []string{"metal", "plants", "chemicals"},
	"properties": bson.M{
		"metal":     bson.M{"bsonType": "int"},
		"plants":    bson.M{"bsonType": "int"},
		"chemicals": bson.M{"bsonType": "int"},
	},
}

var StartingResources = Resources{0, 10, 10}
