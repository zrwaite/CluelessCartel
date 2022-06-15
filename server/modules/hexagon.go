package modules

import (
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/database/models"
)

func GetIndexes(base *models.Base, res *apiModels.Response, X int, Y int) (XIndex int, YIndex int) {
	YIndex = -base.HexagonRows[0].Y + Y
	XIndex = -base.HexagonRows[0].Hexagons[0].X + X

	// Validate range
	if YIndex < 0 || YIndex >= len(base.HexagonRows) {
		res.Errors = append(res.Errors, "Invalid YIndex")
		return
	}
	if XIndex < 0 || XIndex >= len(base.HexagonRows[0].Hexagons) {
		res.Errors = append(res.Errors, "Invalid XIndex")
		return
	}
	return
}
