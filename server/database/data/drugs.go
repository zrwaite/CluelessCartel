package data

import "clueless-cartel-server/database/models"

var Weed = makeRawResource("Weed", 16)

var Fentanyl = models.Resource{
	Name: "Fentanyl",
	SynthesizationCost: []models.ResourcesAmount{
		{NitricAcid.Name, 2},
		{Piperine.Name, 2},
	},
	BatchAmount:        20,
	StartingUnitPrice:  50,
	SynthesizationTime: 45,
}

var Meth = models.Resource{
	Name: "Meth",
	SynthesizationCost: []models.ResourcesAmount{
		{ColdMedecine.Name, 6},
		{Sulfur.Name, 4},
		{RedPhosphorus.Name, 4},
		{Acetone.Name, 5},
		{Lithium.Name, 3},
		{HydrochloricAcid.Name, 5},
	},
	BatchAmount:        40,
	StartingUnitPrice:  20,
	SynthesizationTime: 60,
}

var AllDrugs = []models.Resource{Weed, Fentanyl, Meth}

var StartingDrugs = []models.ResourcesAmount{
	{Weed.Name, 10},
	{Fentanyl.Name, 10},
	{Meth.Name, 10},
}
