package data

import "clueless-cartel-server/database/models"

var Gun = models.Resource{
	Name:               "Gun",
	BatchAmount:        1,
	StartingUnitPrice:  100,
	SynthesizationTime: 100,
	Type:               Solid,
}.AddSynthesizationCost([]models.ResourceAmount{
	makeResourceAmount(Steel, 10),
	makeResourceAmount(Aluminium, 1),
})

var Explosive = models.Resource{
	Name:               "Explosive",
	BatchAmount:        1,
	StartingUnitPrice:  100,
	SynthesizationTime: 100,
	Type:               Solid,
}.AddSynthesizationCost([]models.ResourceAmount{
	makeResourceAmount(Steel, 2),
	makeResourceAmount(Brass, 2),
	makeResourceAmount(GunPowder, 6),
})

var Ammunition = models.Resource{
	Name:               "Ammunition",
	BatchAmount:        10,
	StartingUnitPrice:  1,
	SynthesizationTime: 1,
	Type:               Solid,
}.AddSynthesizationCost([]models.ResourceAmount{
	makeResourceAmount(Brass, 1),
	makeResourceAmount(GunPowder, 1),
})

var StartingWeapons = []models.ResourceAmount{
	makeResourceAmount(Gun, 1),
	makeResourceAmount(Explosive, 1),
	makeResourceAmount(Ammunition, 10),
}

var AllWeapons = []models.Resource{Gun, Explosive, Ammunition}
