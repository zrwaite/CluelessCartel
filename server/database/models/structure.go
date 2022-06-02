package models

import "go.mongodb.org/mongo-driver/bson"

type Structure struct {
	Moveable         bool
	Name             string
	Image            string
	LandMaterials    []string  `bson:"land_materials"`
	ResourceCapacity Resources `bson:"resource_capacity"`
}

var structureSchema = bson.M{
	"bsonType": "object",
	"required": []string{"moveable", "name", "image", "land_materials", "resource_capacity"},
	"properties": bson.M{
		"moveable": bson.M{"bsonType": "bool"},
		"name":     bson.M{"bsonType": "string"},
		"index":    bson.M{"bsonType": "string"},
		"land_materials": bson.M{
			"bsonType":    "array",
			"uniqueItems": true,
			"items":       bson.M{"bsonType": "string"},
		},
		"resource_capacity": resourcesSchema,
	},
}

var EmptyStructure = Structure{
	Moveable:         true,
	Name:             "Empty",
	Image:            "",
	LandMaterials:    []string{},
	ResourceCapacity: Resources{},
}
var RV = Structure{
	Moveable:      true,
	Name:          "RV",
	Image:         "CC.svg",
	LandMaterials: []string{},
	ResourceCapacity: Resources{
		Metal: 7,
	},
}
var BuriedStorage = Structure{
	Moveable:      false,
	Name:          "Buried Storage",
	Image:         "CC.svg",
	LandMaterials: []string{LandMaterial.Dirt, LandMaterial.Grass, LandMaterial.Sand},
}
var StorageUnit = Structure{
	Moveable:      false,
	Name:          "Storage Unit",
	Image:         "CC.svg",
	LandMaterials: []string{LandMaterial.Pavement},
}
var Shed = Structure{
	Moveable:      true,
	Name:          "Shed",
	Image:         "CC.svg",
	LandMaterials: []string{},
}
var Armory = Structure{
	Moveable:      false,
	Name:          "Armory",
	Image:         "CC.svg",
	LandMaterials: []string{LandMaterial.Pavement},
}
