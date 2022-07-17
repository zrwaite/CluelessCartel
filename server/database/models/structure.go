package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Structure struct {
	Moveable      bool
	Enemy         bool
	Natural       bool
	Name          string
	LandMaterials []LandMaterial `bson:"land_materials"`
	Capacities    []ResourceTypeCapacity
	Resources     []ResourceAmount
	Drugs         []ResourceAmount
	Weapons       []ResourceAmount
}

var structureSchema = bson.M{
	"bsonType": "object",
	"required": []string{"moveable", "enemy", "natural", "name", "land_materials", "capacities", "resources", "drugs", "weapons"},
	"properties": bson.M{
		"moveable":       bsonBoolean,
		"natural":        bsonBoolean,
		"enemy":          bsonBoolean,
		"name":           bsonString,
		"land_materials": bsonArray(landMaterialSchema),
		"capacities":     bsonArray(resourceTypeCapacitySchema),
		"resources":      bsonArray(resourcesAmountSchema),
		"drugs":          bsonArray(resourcesAmountSchema),
		"weapons":        bsonArray(resourcesAmountSchema),
	},
}
