package structure

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/api/base"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"clueless-cartel-server/modules"
	"encoding/json"
)

type DriveByStructureParams struct {
	Username     string
	BaseLocation string
	HexagonX     int
	HexagonY     int
}

func handleDriveByStructure(body []byte, res *apiModels.Response) {
	var driveByStructureParams DriveByStructureParams
	baseReader := bytes.NewReader(body)
	err := json.NewDecoder(baseReader).Decode(&driveByStructureParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		models.ValidateData(driveByStructureParams, res)
	}
	if len(res.Errors) != 0 {
		res.Response = driveByStructureParams
		return
	}
	var user *models.GetUserReturn
	user, res.Status = dbModules.GetUserGameData(driveByStructureParams.Username)
	if res.Status == 404 {
		res.Errors = append(res.Errors, "User not found")
		return
	} else if res.Status == 400 {
		res.Errors = append(res.Errors, "Failed to get user")
		return
	} else {
		res.Status = 400
	}

	location, success := models.GetLocation(driveByStructureParams.BaseLocation)
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

	XIndex, YIndex := modules.GetIndexes(base, res, driveByStructureParams.HexagonX, driveByStructureParams.HexagonY)
	if len(res.Errors) != 0 {
		return
	}

	hexagon := &base.HexagonRows[YIndex].Hexagons[XIndex]

	if !driveByStructure(base, hexagon, res) {
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

func getDriveByCost() int {
	return 1
}

func driveByStructure(base *models.Base, hexagon *models.Hexagon, res *apiModels.Response) (success bool) {
	success = false
	if hexagon.Structure.Name == "Empty" {
		res.Errors = append(res.Errors, "Nobody to shoot up")
		return
	}
	/*
		driveByCost := getDriveByCost()
		if base.Weapons.Guns < driveByCost {
			res.Errors = append(res.Errors, "Not enough ammunition")
			return
		}
		base.Weapons.Guns -= driveByCost
	*/
	hexagon.Structure.Enemy = false
	return true
}
