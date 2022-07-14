package data

import "clueless-cartel-server/database/models"

func GetGameData() (gameData *models.Game) {
	gameData = new(models.Game)
	gameData.AllDrugs = AllDrugs
	gameData.AllLocations = AllLocations
	gameData.AllLandMaterials = AllLandMaterials
	gameData.AllResources = []models.Resource{Metal, Plants, Chemicals}
	gameData.AllStructures = AllStructures
	return
}
