package models

import "math/rand"

type Location struct {
	Name         string
	LandMaterial LandMaterial
}

var Nevada = Location{"Nevada", Sand}
var Florida = Location{"Florida", Sand}
var Minnesota = Location{"Minnesota", Grass}
var NewYork = Location{"New York", Pavement}

var AllLocations = []Location{Nevada, Florida, Minnesota, NewYork}

func GetLocation(locationName string) (location Location, success bool) {
	for _, location := range AllLocations {
		if location.Name == locationName {
			return location, true
		}
	}
	success = false
	return
}

func GetSemiRandomLandMaterial(location Location, surroundingHexagons []Hexagon) LandMaterial {
	probType := rand.Intn(4)
	if probType == 0 {
		index := rand.Intn(len(AllLandMaterials))
		return AllLandMaterials[index]
	} else {
		index := rand.Intn(len(surroundingHexagons))
		return surroundingHexagons[index].LandMaterial
	}
}
