package base

import (
	"clueless-cartel-server/database/models"
	"math"
)

func GetBase(userBases *[]models.Base, location string) (base *models.Base, found bool) {
	for _, userBase := range *userBases {
		// Check all bases for if location is in use
		if userBase.Location == location {
			return &userBase, true
		}
	}
	return nil, false
}

func GetBaseCost(bases *[]models.Base) (num int) {
	// 10000 * 10^len(bases)
	return int(10000 * math.Pow(10, float64(len(*bases))))
}
