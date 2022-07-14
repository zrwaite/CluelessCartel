package data

import "clueless-cartel-server/database/models"

var Metal = models.Resource{
	Name: "Metal",
}

var Plants = models.Resource{
	Name: "Plants",
}

var Chemicals = models.Resource{
	Name: "Chemicals",
}

var StartingResources = []models.ResourcesAmount{
	{ResourceName: "Metal", Amount: 10},
	{ResourceName: "Plants", Amount: 10},
	{ResourceName: "Chemicals", Amount: 10},
}
