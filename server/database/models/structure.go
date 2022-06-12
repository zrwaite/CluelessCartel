package models

import "go.mongodb.org/mongo-driver/bson"

type Structure struct {
	Moveable         bool
	Name             string
	LandMaterials    []LandMaterial `bson:"land_materials"`
	ResourceCapacity Resources      `bson:"resource_capacity"`
}

var structureSchema = bson.M{
	"bsonType": "object",
	"required": []string{"moveable", "name", "land_materials", "resource_capacity"},
	"properties": bson.M{
		"moveable": bson.M{"bsonType": "bool"},
		"name":     bson.M{"bsonType": "string"},
		"land_materials": bson.M{
			"bsonType":    "array",
			"uniqueItems": true,
			"items":       landMaterialScema,
		},
		"resource_capacity": resourcesSchema,
	},
}

var EmptyStructure = Structure{
	Moveable:         true,
	Name:             "Empty",
	LandMaterials:    []LandMaterial{},
	ResourceCapacity: Resources{},
}
var RV = Structure{
	Moveable:      true,
	Name:          "RV",
	LandMaterials: []LandMaterial{},
	ResourceCapacity: Resources{
		Metal: 7,
	},
}
var BuriedStorage = Structure{
	Moveable:      false,
	Name:          "Buried Storage",
	LandMaterials: []LandMaterial{Dirt, Grass, Sand},
}
var StorageUnit = Structure{
	Moveable:      false,
	Name:          "Storage Unit",
	LandMaterials: []LandMaterial{Pavement},
}
var Shed = Structure{
	Moveable:      true,
	Name:          "Shed",
	LandMaterials: []LandMaterial{},
}
var Armory = Structure{
	Moveable:      false,
	Name:          "Armory",
	LandMaterials: []LandMaterial{Pavement},
}

var AllStructures = []Structure{EmptyStructure, RV, BuriedStorage, StorageUnit, Shed, Armory}

func GetStructure(structureName string) (structure Structure, success bool) {
	for _, structure := range AllStructures {
		if structure.Name == structureName {
			return structure, true
		}
	}
	success = false
	return
}
