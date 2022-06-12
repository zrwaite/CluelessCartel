package models

type Game struct {
	//All base locations
	AllDrugs         []Drug
	AllLocations     []Location
	AllResources     []Resource
	AllStructures    []Structure
	AllLandMaterials []LandMaterial
}

func GetGameData() (gameData *Game) {
	gameData = new(Game)
	gameData.AllDrugs = []Drug{Opioids, Weed, Meth}
	gameData.AllLocations = AllLocations
	gameData.AllLandMaterials = AllLandMaterials
	gameData.AllResources = []Resource{Metal, Plants, Chemicals}
	gameData.AllStructures = AllStructures
	return
}
