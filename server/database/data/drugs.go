package data

import "clueless-cartel-server/database/models"

var Weed = makeRawResource("Weed", 16, Solid)

var Fentanyl = models.Resource{
	Name:               "Fentanyl",
	BatchAmount:        20,
	StartingUnitPrice:  50,
	SynthesizationTime: 45,
	Type:               Powder,
}.AddSynthesizationCost([]models.ResourceAmount{
	makeResourceAmount(NitricAcid, 2),
	makeResourceAmount(Piperine, 2),
})

var Meth = models.Resource{
	Name:               "Meth",
	BatchAmount:        40,
	StartingUnitPrice:  20,
	SynthesizationTime: 60,
	Type:               Powder,
}.AddSynthesizationCost([]models.ResourceAmount{
	makeResourceAmount(ColdMedecine, 6),
	makeResourceAmount(Sulfur, 4),
	makeResourceAmount(RedPhosphorus, 4),
	makeResourceAmount(Acetone, 5),
	makeResourceAmount(Lithium, 3),
	makeResourceAmount(HydrochloricAcid, 5),
})

var AllDrugs = []models.Resource{Weed, Fentanyl, Meth}

var StartingDrugs = []models.ResourceAmount{
	makeResourceAmount(Weed, 10),
	makeResourceAmount(Fentanyl, 10),
	makeResourceAmount(Meth, 10),
}
