package hexagon

import (
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/database/models"
	"math"
)

func GetNumHexagons(hexagonRows *[]models.HexagonRow) (num int) {
	for _, hexagonRow := range *hexagonRows {
		num += len(hexagonRow.Hexagons)
	}
	return
}

func GetHexagonCost(base *models.Base) (num int) {
	// 1000 * 10^index * 1.5^len(hexagons)
	return 1000 * int(math.Pow(10, float64(base.Index))*math.Pow(1.5, float64(GetNumHexagons(&base.HexagonRows))))
}

func BuyHexagon(hexagonRows *[]models.HexagonRow, HexagonX int, HexagonY int, res *apiModels.Response) (success bool) {

	return false
}
