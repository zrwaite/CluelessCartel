package data

import (
	"clueless-cartel-server/database/models"
	"errors"
	"math/rand"
)

func GetRandomLandMaterial() (material string) {
	materials := []string{"Sand", "Grass", "Dirt", "Pavement"}
	return materials[rand.Intn(len(materials))]
}

func CreateStartingBase(locationName string, index int) (models.Base, error) {
	var base models.Base
	location, success := GetLocation(locationName)
	if !success {
		return base, errors.New("invalid location")
	}
	base = models.Base{
		Location:    location,
		Index:       index,
		Resources:   StartingResources,
		Drugs:       StartingDrugs,
		Weapons:     StartingWeapons,
		HexagonRows: CreateStartingHexagonRows(location),
	}
	return base, nil

}
