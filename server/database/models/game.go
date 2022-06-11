package models

type Game struct {
	//All base locations
	AllDrugs     []Drug
	AllLocations []Location
	AllResources []Resource
}

func GetGameData() (gameData *Game) {
	gameData = new(Game)
	gameData.AllDrugs = []Drug{Opioids, Weed, Meth}
	gameData.AllLocations = []Location{Nevada, Florida, CanadianBorder, NewYork}
	gameData.AllResources = []Resource{Metal, Plants, Chemicals}
	return
}
