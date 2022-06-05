package hexagon

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/api/base"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"encoding/json"
)

type BuyHexagonParams struct {
	Username     string
	BaseLocation string
	HexagonX     int
	HexagonY     int
}

func buyHexagon(body []byte, res *apiModels.Response) {
	var buyHexagonParams BuyHexagonParams
	baseReader := bytes.NewReader(body)
	err := json.NewDecoder(baseReader).Decode(&buyHexagonParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		models.ValidateData(buyHexagonParams, res)
	}
	if len(res.Errors) == 0 {
		var user *models.GetUserReturn
		user, res.Status = dbModules.GetUserGameData(buyHexagonParams.Username)
		if res.Status == 404 {
			res.Errors = append(res.Errors, "User not found")
			return
		} else if res.Status == 400 {
			res.Errors = append(res.Errors, "Failed to get user")
			return
		} else {
			res.Status = 400
		}

		// Make sure base exists at location
		base, baseExists := base.GetBase(&user.Bases, buyHexagonParams.BaseLocation)
		if !baseExists {
			res.Errors = append(res.Errors, "Base doesn't exist at that location")
			return
		}

		//Find the cost for the new hexagon
		hexagonCost := GetHexagonCost(base)
		if user.Cash < hexagonCost {
			res.Errors = append(res.Errors, "Not enough cash!")
			return
		}

		// Update user cash, and add base
		user.Cash -= hexagonCost

		// Attempt to buy hexagon, update user on success
		if !BuyHexagon(&base.HexagonRows, buyHexagonParams.HexagonX, buyHexagonParams.HexagonY, res) {
			return
		}
		if !models.ValidateHexagonRowsStructure(base.HexagonRows) {
			res.Errors = append(res.Errors, "Failed to create base")
			return
		}
		success := dbModules.UpdateUser(user)
		if !success {
			res.Errors = append(res.Errors, "Failed to update user - "+err.Error())
			return
		}
		res.Success = true
		res.Status = 200
		res.Response = user
	} else {
		res.Response = buyHexagonParams
	}
}
