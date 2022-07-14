package data

import "clueless-cartel-server/database/models"

func GetGameData() (gameData *models.Game) {
	gameData = new(models.Game)
	gameData.AllDrugs = AllDrugs
	gameData.AllLocations = AllLocations
	gameData.AllLandMaterials = AllLandMaterials
	gameData.AllResources = AllResources
	gameData.AllStructures = AllStructures
	return
}
