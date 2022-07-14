package reset

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/database/data"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"encoding/json"
	"net/http"
)

func ResetHandler(r *http.Request, body []byte, res *apiModels.Response) {
	if r.Method != "DELETE" {
		res.Errors = append(res.Errors, "Method "+r.Method+" is not supported")
		return
	}
	var userParams models.GetUsernameStruct
	userReader := bytes.NewReader(body)
	err := json.NewDecoder(userReader).Decode(&userParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		data.ValidateData(userParams, res)
	}
	if len(res.Errors) != 0 {
		return
	}
	var user *models.GetUserReturn
	user, res.Status = dbModules.GetUserGameData(userParams.Username)
	if res.Status == 404 {
		res.Errors = append(res.Errors, "User not found")
		return
	} else if res.Status == 400 {
		res.Errors = append(res.Errors, "Error while finding user")
		return
	}

	user.Cash = models.STARTING_USER_CASH
	user.Bases = []models.Base{}

	success := dbModules.UpdateUser(user)
	if !success {
		res.Errors = append(res.Errors, "Failed to update user - "+err.Error())
		return
	}
	res.Success = true
	res.Status = 200
	res.Response = user
}
