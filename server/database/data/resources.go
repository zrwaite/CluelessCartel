package data

import "clueless-cartel-server/database/models"

// Raw resources
var ColdMedecine = makeRawResource("Cold Medecine", 20)
var HydrochloricAcid = makeRawResource("Hydrochloric Acid", 5)
var NitricAcid = makeRawResource("Nitric Acid", 4)
var Lithium = makeRawResource("Lithium", 5)
var Aluminium = makeRawResource("Aluminium", 3)
var Pepper = makeRawResource("Pepper", 1)
var Metal = makeRawResource("Metal", 6)
var Wood = makeRawResource("Wood", 2)
var Nitrate = makeRawResource("Nitrate", 1)
var Sulfur = makeRawResource("Sulfur", 2)
var RedPhosphorus = makeRawResource("Red Phosphorus", 2)
var Acetone = makeRawResource("Acetone", 2)
var Ethanol = makeRawResource("Ethanol", 2)

var Piperine = models.Resource{
	Name:              "Piperine",
	StartingUnitPrice: 10,
	SynthesizationCost: []models.ResourcesAmount{
		{Ethanol.Name, 1},
		{Charcoal.Name, 1},
		{Pepper.Name, 1},
		{Aluminium.Name, 1},
	},
	SynthesizationTime: 30,
}

var Charcoal = models.Resource{
	Name:               "Charcoal",
	SynthesizationCost: []models.ResourcesAmount{{Wood.Name, 1}},
	BatchAmount:        1,
	StartingUnitPrice:  1,
	SynthesizationTime: 10,
}

var GunPowder = models.Resource{
	Name:               "GunPowder",
	SynthesizationCost: []models.ResourcesAmount{{Nitrate.Name, 15}, {Charcoal.Name, 3}, {Sulfur.Name, 2}},
	StartingUnitPrice:  25,
	SynthesizationTime: 5,
}

func makeRawResource(name string, price int) models.Resource {
	return models.Resource{
		Name:              name,
		StartingUnitPrice: price,
	}
}

var StartingResourceAmounts = []models.ResourcesAmount{
	{ResourceName: Metal.Name, Amount: 10},
	{ResourceName: "", Amount: 10},
}

var AllResources = []models.Resource{
	Metal,
	Nitrate,
	GunPowder,
	Wood,
	Charcoal,
	Sulfur,
	RedPhosphorus,
	Acetone,
	ColdMedecine,
	HydrochloricAcid,
	Lithium,
	Aluminium,
	Pepper,
	Piperine,
}
