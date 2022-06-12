package models

type Location struct {
	Name         string
	LandMaterial LandMaterial
}

var Nevada = Location{"Nevada", Sand}
var Florida = Location{"Florida", Sand}
var Minnesota = Location{"Minnesota", Grass}
var NewYork = Location{"New York", Pavement}

func GetLocation(locationName string) (location Location, success bool) {
	success = true
	switch locationName {
	case "Nevada":
		location = Nevada
	case "Florida":
		location = Florida
	case "Minnesota":
		location = Minnesota
	case "New York":
		location = NewYork
	default:
		success = false
	}
	return
}
