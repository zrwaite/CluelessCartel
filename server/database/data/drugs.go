package data

import "clueless-cartel-server/database/models"

var Opioids = models.Drug{
	Cost:        []models.ResourceCost{{"Plants", 5}, {"Chemicals", 5}},
	BatchAmount: 20,
	UnitPrice:   50,
}

var Weed = models.Drug{
	Cost:        []models.ResourceCost{{"Plants", 10}},
	BatchAmount: 50,
	UnitPrice:   16,
}

var Meth = models.Drug{
	Cost:        []models.ResourceCost{{"Chemicals", 10}},
	BatchAmount: 40,
	UnitPrice:   20,
}

var AllDrugs = []models.Drug{Opioids, Weed, Meth}

var StartingDrugs = []models.DrugAmount{
	{"Opioids", 10},
	{"Weed", 10},
	{"Meth", 10},
}
