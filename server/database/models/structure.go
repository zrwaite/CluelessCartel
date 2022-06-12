package models

import (
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
)

type Structure struct {
	Moveable         bool
	Enemy            bool
	Natural          bool
	Name             string
	LandMaterials    []LandMaterial `bson:"land_materials"`
	ResourceCapacity Resources      `bson:"resource_capacity"`
}

var structureSchema = bson.M{
	"bsonType": "object",
	"required": []string{"moveable", "enemy", "natural", "name", "land_materials", "resource_capacity"},
	"properties": bson.M{
		"moveable": bson.M{"bsonType": "bool"},
		"natural":  bson.M{"bsonType": "bool"},
		"enemy":    bson.M{"bsonType": "bool"},
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
	Moveable:      true,
	Name:          "Empty",
	LandMaterials: []LandMaterial{},
}
var RV = Structure{
	Moveable: true,
	Name:     "RV",
	ResourceCapacity: Resources{
		Metal: 7,
	},
	LandMaterials: []LandMaterial{Dirt, Grass, Sand, Pavement},
}
var BuriedStorage = Structure{
	Name:          "Buried Storage",
	LandMaterials: []LandMaterial{Dirt, Grass, Sand},
}
var StorageUnit = Structure{
	Name:          "Storage Unit",
	LandMaterials: []LandMaterial{Pavement},
}
var Shed = Structure{
	Moveable:      true,
	Name:          "Shed",
	LandMaterials: []LandMaterial{Dirt, Grass, Sand, Pavement},
}
var Armory = Structure{
	Name:          "Armory",
	LandMaterials: []LandMaterial{Pavement},
}

var Trees = Structure{
	Name:          "Trees",
	LandMaterials: []LandMaterial{Dirt, Grass},
	Natural:       true,
}

var Rocks = Structure{
	Name:          "Rocks",
	LandMaterials: []LandMaterial{Dirt, Grass, Sand},
	Natural:       true,
}

var Garbage = Structure{
	Name:          "Garbage",
	LandMaterials: []LandMaterial{Dirt, Pavement, Water},
	Natural:       true,
}

var NaturalStructure = []Structure{Trees, Rocks, Garbage}

var AllStructures = append([]Structure{EmptyStructure, RV, BuriedStorage, StorageUnit, Shed, Armory}, NaturalStructure...)

func GetStructure(structureName string) (structure Structure, success bool) {
	for _, structure := range AllStructures {
		if structure.Name == structureName {
			return structure, true
		}
	}
	success = false
	return
}

func GetSemiRandomStructure(landMaterial LandMaterial) (structure Structure) {
	structure = EmptyStructure
	probability := rand.Intn(5)
	if probability != 0 {
		return
	}
	var referenceStructures []Structure
	if probability == 1 {
		referenceStructures = AllStructures
	} else {
		referenceStructures = NaturalStructure
	}

	structureList := []Structure{}
	for _, listStructure := range referenceStructures {
		valid := false
		if len(listStructure.LandMaterials) == 0 {
			valid = true
		} else {
			for _, validMaterial := range listStructure.LandMaterials {
				if landMaterial == validMaterial {
					valid = true
				}
			}
		}
		if valid {
			structureList = append(structureList, listStructure)
		}
	}
	if len(structureList) != 0 {
		index := rand.Intn(len(structureList))
		structure = structureList[index]
		if structure.Name != "Empty" && !structure.Natural {
			structure.Enemy = true
		}
	}
	return
}
