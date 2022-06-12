package structure

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/api/base"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"encoding/json"
)

type RemoveStructureParams struct {
	Username     string
	BaseLocation string
	HexagonX     int
	HexagonY     int
}

func removeStructure(body []byte, res *apiModels.Response) {
	var removeStructureParams RemoveStructureParams
	baseReader := bytes.NewReader(body)
	err := json.NewDecoder(baseReader).Decode(&removeStructureParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		models.ValidateData(removeStructureParams, res)
	}
	if len(res.Errors) != 0 {
		res.Response = removeStructureParams
		return
	}
	var user *models.GetUserReturn
	user, res.Status = dbModules.GetUserGameData(removeStructureParams.Username)
	if res.Status == 404 {
		res.Errors = append(res.Errors, "User not found")
		return
	} else if res.Status == 400 {
		res.Errors = append(res.Errors, "Failed to get user")
		return
	} else {
		res.Status = 400
	}

	location, success := models.GetLocation(removeStructureParams.BaseLocation)
	if !success {
		res.Errors = append(res.Errors, "invalid location")
		return
	}

	// Make sure base exists at location
	base, baseExists := base.GetBase(&user.Bases, location)
	if !baseExists {
		res.Errors = append(res.Errors, "Base doesn't exist at that location")
		return
	}

	YIndex := -base.HexagonRows[0].Y + removeStructureParams.HexagonY
	XIndex := -base.HexagonRows[0].Hexagons[0].X + removeStructureParams.HexagonX

	// Validate range
	if YIndex < 0 || YIndex >= len(base.HexagonRows) {
		res.Errors = append(res.Errors, "Invalid YIndex")
		return
	}
	if XIndex < 0 || XIndex >= len(base.HexagonRows[0].Hexagons) {
		res.Errors = append(res.Errors, "Invalid XIndex")
		return
	}

	hexagon := &base.HexagonRows[YIndex].Hexagons[XIndex]

	/*
		//Find the cost for the new structure
		structureCost := GetStructureCost(base)
		user.Cash += hexagonCost
		if user.Cash < hexagonCost {
			res.Errors = append(res.Errors, "Not enough cash!")
			return
		}

		// Update user cash, and add base
		user.Cash -= hexagonCost
	*/
	if !RemoveStructure(hexagon, res) {
		return
	}
	success = dbModules.UpdateUser(user)
	if !success {
		res.Errors = append(res.Errors, "Failed to update user - "+err.Error())
		return
	}
	res.Success = true
	res.Status = 200
	res.Response = user
}

func RemoveStructure(hexagon *models.Hexagon, res *apiModels.Response) (success bool) {
	success = false
	if !hexagon.Owned {
		res.Errors = append(res.Errors, "Hexagon not owned")
		return
	}
	if hexagon.Structure.Name == "Empty" {
		res.Errors = append(res.Errors, "Nothing to clear")
		return
	}
	hexagon.Structure = models.EmptyStructure
	return true
}
