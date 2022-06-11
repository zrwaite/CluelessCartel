package base

import (
	"clueless-cartel-server/database/models"
	"math"
)

func GetBase(userBases *[]models.Base, location models.Location) (base *models.Base, found bool) {
	for i := 0; i < len(*userBases); i++ {
		// Check all bases for if location is in use
		if (*userBases)[i].Location.Name == location.Name {
			return &(*userBases)[i], true
		}
	}
	return nil, false
}

func GetBaseCost(bases *[]models.Base) (num int) {
	// 10000 * 10^len(bases)
	return int(10000 * math.Pow(10, float64(len(*bases))))
}
