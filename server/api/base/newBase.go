package base

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"encoding/json"
	"math"
)

type PostBaseParams struct {
	Username string
	Location string
}

func newBase(body []byte, res *apiModels.Response) {
	var newBaseParams PostBaseParams
	baseReader := bytes.NewReader(body)
	err := json.NewDecoder(baseReader).Decode(&newBaseParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		models.ValidateData(newBaseParams, res)
	}
	if len(res.Errors) == 0 {
		var user *models.GetUserReturn
		base, err := models.CreateStartingBase(newBaseParams.Location)
		if err != nil {
			res.Errors = append(res.Errors, err.Error())
			return
		}
		base.Location = newBaseParams.Location
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
		for _, userBase := range user.Bases {
			// Check all bases for if location is in use
			if userBase.Location == base.Location {
				res.Errors = append(res.Errors, "Location already used!")
				return
			}
		}
		//Set new base location
		//Find the cost for the new base
		baseCost := int(10000 * math.Pow(10, float64(len(user.Bases))))
		if user.Cash < baseCost {
			res.Errors = append(res.Errors, "Not enough cash!")
			return
		}
		// Update user cash, and add base
		user.Cash -= baseCost
		user.Bases = append(user.Bases, base)
		success := dbModules.UpdateUser(user)
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
