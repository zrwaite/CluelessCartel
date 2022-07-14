package base

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/database/data"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"encoding/json"
)

type NewBaseParams struct {
	Username string
	Location string
}

func handleNewBase(body []byte, res *apiModels.Response) {
	var newBaseParams NewBaseParams
	baseReader := bytes.NewReader(body)
	err := json.NewDecoder(baseReader).Decode(&newBaseParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		data.ValidateData(newBaseParams, res)
	}
	if len(res.Errors) == 0 {
		var user *models.GetUserReturn
		user, res.Status = dbModules.GetUserGameData(newBaseParams.Username)
		if res.Status == 404 {
			res.Errors = append(res.Errors, "User not found")
			return
		} else if res.Status == 400 {
			res.Errors = append(res.Errors, "Failed to get user")
			return
		} else {
			res.Status = 400
		}
		base, err := data.CreateStartingBase(newBaseParams.Location, len(user.Bases))
		if err != nil {
			res.Errors = append(res.Errors, err.Error())
			return
		}
		var success bool
		base.Location, success = data.GetLocation(newBaseParams.Location)
		if !success {
			res.Errors = append(res.Errors, "invalid location")
			return
		}
		_, baseFound := GetBase(&user.Bases, base.Location)
		if baseFound {
			// If location is already in use, a new base can not be created there
			res.Errors = append(res.Errors, "Location already used!")
			return
		}

		/*
			//Find the cost for the new base
			baseCost := GetBaseCost(&user.Bases)
			if user.Cash < baseCost {
				res.Errors = append(res.Errors, "Not enough cash!")
				return
			}

			// Update user cash, and add base
			user.Cash -= baseCost
		*/

		user.Bases = append(user.Bases, base)
		if !data.ValidateHexagonRowsStructure(&base.HexagonRows) {
			res.Errors = append(res.Errors, "Failed to create base")
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
	} else {
		res.Response = newBaseParams
	}
}
