package structure

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/api/base"
	"clueless-cartel-server/database/data"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"clueless-cartel-server/modules"
	"encoding/json"
)

type ExplodeStructureParams struct {
	Username     string
	BaseLocation string
	HexagonX     int
	HexagonY     int
}

func handleExplodeStructure(body []byte, res *apiModels.Response) {
	var explodeStructureParams ExplodeStructureParams
	baseReader := bytes.NewReader(body)
	err := json.NewDecoder(baseReader).Decode(&explodeStructureParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		data.ValidateData(explodeStructureParams, res)
	}
	if len(res.Errors) != 0 {
		res.Response = explodeStructureParams
		return
	}
	var user *models.GetUserReturn
	user, res.Status = dbModules.GetUserGameData(explodeStructureParams.Username)
	if res.Status == 404 {
		res.Errors = append(res.Errors, "User not found")
		return
	} else if res.Status == 400 {
		res.Errors = append(res.Errors, "Failed to get user")
		return
	} else {
		res.Status = 400
	}

	location, success := data.GetLocation(explodeStructureParams.BaseLocation)
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

	XIndex, YIndex := modules.GetIndexes(base, res, explodeStructureParams.HexagonX, explodeStructureParams.HexagonY)
	if len(res.Errors) != 0 {
		return
	}

	hexagon := &base.HexagonRows[YIndex].Hexagons[XIndex]

	if !explodeStructure(base, hexagon, res) {
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

func getExplosionCost() int {
	return 1
}

func explodeStructure(base *models.Base, hexagon *models.Hexagon, res *apiModels.Response) (success bool) {
	success = false
	if hexagon.Structure.Name == "Empty" {
		res.Errors = append(res.Errors, "Nothing to clear")
		return
	}
	/*
		explodeCost := getExplosionCost()
		if base.Weapons.Explosives < explodeCost {
			res.Errors = append(res.Errors, "Not enough explosives")
			return
		}
		base.Weapons.Explosives -= explodeCost
	*/
	hexagon.Structure = data.EmptyStructure
	return true
}
