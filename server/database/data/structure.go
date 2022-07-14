package data

import (
	"clueless-cartel-server/database/models"
	"math/rand"
)

var EmptyStructure = models.Structure{
	Moveable:      true,
	Name:          "Empty",
	LandMaterials: []models.LandMaterial{},
}
var RV = models.Structure{
	Moveable: true,
	Name:     "RV",
	ResourceCapacity: []models.ResourcesAmount{
		{ResourceName: "Metal", Amount: 1},
	},
	LandMaterials: []models.LandMaterial{Dirt, Grass, Sand, Pavement},
}
var BuriedStorage = models.Structure{
	Name:          "Buried Storage",
	LandMaterials: []models.LandMaterial{Dirt, Grass, Sand},
}
var StorageUnit = models.Structure{
	Name:          "Storage Unit",
	LandMaterials: []models.LandMaterial{Pavement},
}
var Shed = models.Structure{
	Moveable:      true,
	Name:          "Shed",
	LandMaterials: []models.LandMaterial{Dirt, Grass, Sand, Pavement},
}

var Farm = models.Structure{
	Name:          "Farm",
	LandMaterials: []models.LandMaterial{Dirt},
}

var Armory = models.Structure{
	Name:          "Armory",
	LandMaterials: []models.LandMaterial{Pavement},
}

var Trees = models.Structure{
	Name:          "Trees",
	LandMaterials: []models.LandMaterial{Dirt, Grass},
	Natural:       true,
}

var Rocks = models.Structure{
	Name:          "Rocks",
	LandMaterials: []models.LandMaterial{Dirt, Grass, Sand},
	Natural:       true,
}

var Garbage = models.Structure{
	Name:          "Garbage",
	LandMaterials: []models.LandMaterial{Dirt, Pavement, Water},
	Natural:       true,
}

var NaturalStructures = []models.Structure{Trees, Rocks, Garbage}
var ManMadeStructures = []models.Structure{RV, BuriedStorage, StorageUnit, Shed, Armory, Farm}

var AllStructures = append(append(ManMadeStructures, NaturalStructures...), EmptyStructure)

func GetStructure(structureName string) (structure models.Structure, success bool) {
	for _, structure := range AllStructures {
		if structure.Name == structureName {
			return structure, true
		}
	}
	success = false
	return
}

func GetSemiRandomStructure(landMaterial models.LandMaterial, enemyCamp bool, forest bool) (structure models.Structure) {
	structure = EmptyStructure
	var referenceStructures []models.Structure
	probability := rand.Intn(40)
	if enemyCamp {
		if probability < 25 {
			return
		} else if probability > 25 && probability < 30 {
			referenceStructures = NaturalStructures
		} else {
			referenceStructures = ManMadeStructures
		}
	} else if forest {
		if probability < 3 {
			return
		} else if probability == 36 || probability == 37 {
			referenceStructures = ManMadeStructures
		} else if probability > 37 {
			referenceStructures = NaturalStructures
		} else {
			referenceStructures = []models.Structure{Trees}
		}
	} else {
		if probability < 30 {
			return
		}
		if probability < 32 {
			referenceStructures = ManMadeStructures
		} else {
			referenceStructures = append(NaturalStructures, []models.Structure{Trees, Trees}...)
		}
	}

	structureList := []models.Structure{}
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
