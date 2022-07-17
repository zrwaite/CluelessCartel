package data

import (
	"clueless-cartel-server/database/models"
)

var Liquid = models.ResourceType{Name: "Liquid"}
var Solid = models.ResourceType{Name: "Solid"}
var Powder = models.ResourceType{Name: "Powder"}

// Raw resources
var ColdMedecine = makeRawResource("Cold Medecine", 20, Liquid)
var HydrochloricAcid = makeRawResource("Hydrochloric Acid", 5, Liquid)
var NitricAcid = makeRawResource("Nitric Acid", 4, Liquid)
var Lithium = makeRawResource("Lithium", 5, Solid)
var Aluminium = makeRawResource("Aluminium", 3, Solid)
var Pepper = makeRawResource("Pepper", 1, Powder)
var Steel = makeRawResource("Steel", 6, Solid)
var Brass = makeRawResource("Brass", 6, Solid)
var Wood = makeRawResource("Wood", 2, Solid)
var Nitrate = makeRawResource("Nitrate", 1, Powder)
var Sulfur = makeRawResource("Sulfur", 2, Powder)
var RedPhosphorus = makeRawResource("Red Phosphorus", 2, Powder)
var Acetone = makeRawResource("Acetone", 2, Liquid)
var Ethanol = makeRawResource("Ethanol", 2, Liquid)

var Piperine = models.Resource{
	Name:              "Piperine",
	StartingUnitPrice: 10,
	SynthesizationCost: []models.ResourceAmount{
		makeResourceAmount(Ethanol, 1),
		makeResourceAmount(Charcoal, 1),
		makeResourceAmount(Pepper, 1),
		makeResourceAmount(Aluminium, 1),
	},
	SynthesizationTime: 30,
	Type:               Powder,
}

var Charcoal = models.Resource{
	Name: "Charcoal",
	SynthesizationCost: []models.ResourceAmount{
		makeResourceAmount(Wood, 1),
	},
	BatchAmount:        1,
	StartingUnitPrice:  1,
	SynthesizationTime: 10,
	Type:               Powder,
}

var GunPowder = models.Resource{
	Name: "Gun Powder",
	SynthesizationCost: []models.ResourceAmount{
		makeResourceAmount(Nitrate, 15),
		makeResourceAmount(Charcoal, 3),
		makeResourceAmount(Sulfur, 2),
	},
	StartingUnitPrice:  25,
	SynthesizationTime: 5,
	Type:               Powder,
}

func makeRawResource(name string, price int, resourceType models.ResourceType) models.Resource {
	return models.Resource{
		Name:              name,
		StartingUnitPrice: price,
		Type:              resourceType,
	}
}

func makeResourceAmount(resource models.Resource, amount int) models.ResourceAmount {
	return models.ResourceAmount{
		Name:    resource.Name,
		Amount:  amount,
		Quality: 100,
	}
}

func makeResourceTypeCapacity(resourceType models.ResourceType, amount int) models.ResourceTypeCapacity {
	return models.ResourceTypeCapacity{
		Name:     resourceType.Name,
		Capacity: amount,
	}
}

var StartingResourceAmounts = []models.ResourceAmount{
	{Name: Steel.Name, Amount: 10},
	{Name: "", Amount: 10},
}

var AllResources = []models.Resource{
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
	Ethanol,
	Steel,
	Brass,
}
