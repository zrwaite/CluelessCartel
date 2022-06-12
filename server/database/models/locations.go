package models

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
