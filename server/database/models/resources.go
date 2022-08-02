package models

import "go.mongodb.org/mongo-driver/bson"

type Resource struct {
	Name               string
	synthesizationCost []ResourceAmount
	SynthesizationTime int // in minutes
	BatchAmount        int
	StartingUnitPrice  int
	Type               ResourceType
}

func (resource Resource) AddSynthesizationCost(synthesizationCost []ResourceAmount) Resource {
	resource.synthesizationCost = synthesizationCost
	return resource
}

type ResourceType struct {
	Name string
}

type ResourceTypeCapacity struct {
	Name     string
	Capacity int
}

var resourceTypeCapacitySchema = bson.M{
	"bsonType": "object",
	"required": []string{"name", "capacity"},
	"properties": bson.M{
		"name":     bsonString,
		"capacity": bsonInt,
	},
}

type ResourceAmount struct {
	Name    string
	Amount  int
	Quality float64 //percentage
}

var resourcesAmountSchema = bson.M{
	"bsonType": "object",
	"required": []string{"name", "amount", "quality"},
	"properties": bson.M{
		"name":    bsonString,
		"amount":  bsonInt,
		"quality": bsonDouble,
	},
}
