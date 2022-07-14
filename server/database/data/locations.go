package data

import (
	"clueless-cartel-server/database/models"
	"math/rand"
)

var Nevada = models.Location{"Nevada", Sand}
var Florida = models.Location{"Florida", Sand}
var Minnesota = models.Location{"Minnesota", Grass}
var NewYork = models.Location{"New York", Pavement}

var AllLocations = []models.Location{Nevada, Florida, Minnesota, NewYork}

func GetLocation(locationName string) (location models.Location, success bool) {
	for _, location := range AllLocations {
		if location.Name == locationName {
			return location, true
		}
	}
	success = false
	return
}

func GetSemiRandomLandMaterial(location models.Location, surroundingHexagons []models.Hexagon) models.LandMaterial {
	probType := rand.Intn(4)
	if probType == 0 {
		index := rand.Intn(len(AllLandMaterials))
		return AllLandMaterials[index]
	} else {
		index := rand.Intn(len(surroundingHexagons))
		return surroundingHexagons[index].LandMaterial
	}
}
