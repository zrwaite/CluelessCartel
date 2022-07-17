package data

import (
	"clueless-cartel-server/database/models"
	"math/rand"
)

var EmptyStructure = createStructure(
	"Empty",
	[]models.LandMaterial{},
	[]models.ResourceTypeCapacity{},
	false,
)
var RV = createStructure(
	"RV",
	[]models.LandMaterial{Sand, Grass, Dirt, Pavement},
	[]models.ResourceTypeCapacity{
		makeResourceTypeCapacity(Liquid, 10),
		makeResourceTypeCapacity(Solid, 10),
		makeResourceTypeCapacity(Powder, 10),
	},
	true,
)
var BuriedStorage = createStructure(
	"Buried Storage",
	[]models.LandMaterial{Sand, Grass, Dirt},
	[]models.ResourceTypeCapacity{
		makeResourceTypeCapacity(Liquid, 100),
		makeResourceTypeCapacity(Solid, 5),
		makeResourceTypeCapacity(Powder, 10),
	},
	false,
)
var StorageUnit = createStructure(
	"Storage Unit",
	[]models.LandMaterial{Pavement},
	[]models.ResourceTypeCapacity{
		makeResourceTypeCapacity(Liquid, 20),
		makeResourceTypeCapacity(Solid, 60),
		makeResourceTypeCapacity(Powder, 30),
	},
	false,
)
var Shed = createStructure(
	"Shed",
	[]models.LandMaterial{Sand, Grass, Dirt, Pavement},
	[]models.ResourceTypeCapacity{
		makeResourceTypeCapacity(Liquid, 30),
		makeResourceTypeCapacity(Solid, 30),
		makeResourceTypeCapacity(Powder, 30),
	},
	true,
)

var Farm = createStructure(
	"Farm",
	[]models.LandMaterial{Dirt},
	[]models.ResourceTypeCapacity{
		makeResourceTypeCapacity(Liquid, 10),
		makeResourceTypeCapacity(Solid, 20),
		makeResourceTypeCapacity(Powder, 15),
	},
	false,
)

var Armory = createStructure(
	"Armory",
	[]models.LandMaterial{Pavement},
	[]models.ResourceTypeCapacity{
		makeResourceTypeCapacity(Solid, 40),
	},
	false,
)

var Trees = createNaturalStructure(
	"Trees",
	[]models.LandMaterial{Dirt, Grass},
	false,
)

var Rocks = createNaturalStructure(
	"Rocks",
	[]models.LandMaterial{Sand, Dirt, Grass},
	false,
)

var Garbage = createNaturalStructure(
	"Garbage",
	[]models.LandMaterial{Pavement, Dirt, Water},
	false,
)

var Road = createNaturalStructure(
	"Road",
	[]models.LandMaterial{Dirt, Pavement, Grass, Sand},
	false,
)
var NaturalStructures = []models.Structure{Trees, Rocks, Garbage}
var ManMadeStructures = []models.Structure{RV, BuriedStorage, StorageUnit, Shed, Armory, Farm}

var AllStructures = append(append(ManMadeStructures, NaturalStructures...), Road, EmptyStructure)

func createStructure(name string, landMaterials []models.LandMaterial, capacities []models.ResourceTypeCapacity, moveable bool) models.Structure {
	return models.Structure{
		Name:          name,
		LandMaterials: landMaterials,
		Capacities:    capacities,
		Moveable:      moveable,
		Natural:       false,
		Enemy:         false,
		Resources:     []models.ResourceAmount{},
		Drugs:         []models.ResourceAmount{},
		Weapons:       []models.ResourceAmount{},
	}
}

func createNaturalStructure(name string, landMaterials []models.LandMaterial, moveable bool) models.Structure {
	return models.Structure{
		Name:          name,
		LandMaterials: landMaterials,
		Capacities:    []models.ResourceTypeCapacity{},
		Moveable:      moveable,
		Natural:       true,
		Enemy:         false,
		Resources:     []models.ResourceAmount{},
		Drugs:         []models.ResourceAmount{},
		Weapons:       []models.ResourceAmount{},
	}
}

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
