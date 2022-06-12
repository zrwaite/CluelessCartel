package models

type Game struct {
	//All base locations
	AllDrugs      []Drug
	AllLocations  []Location
	AllResources  []Resource
	AllStructures []Structure
}

func GetGameData() (gameData *Game) {
	gameData = new(Game)
	gameData.AllDrugs = []Drug{Opioids, Weed, Meth}
	gameData.AllLocations = []Location{Nevada, Florida, Minnesota, NewYork}
	gameData.AllResources = []Resource{Metal, Plants, Chemicals}
	gameData.AllStructures = AllStructures
	return
}
