package signin

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/api/user"
	"clueless-cartel-server/auth"
	"clueless-cartel-server/auth/tokens"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"encoding/json"
	"net/http"
)

func SignInHandler(r *http.Request, data []byte, res *apiModels.Response) {
	if r.Method != "POST" {
		res.Errors = append(res.Errors, "Method "+r.Method+" is not supported")
		return
	}
	var userParams models.PostUserParams
	userReader := bytes.NewReader(data)
	err := json.NewDecoder(userReader).Decode(&userParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		models.ValidateData(userParams, res)
		if !user.UsernameUsed(userParams.Username) {
			res.Errors = append(res.Errors, "User does not exist")
			res.Status = 404
		}
	}
	if len(res.Errors) != 0 {
		return
	}
	var user *models.GetUserAuthReturn
	user, res.Status = dbModules.GetUserAuthData(userParams.Username)
	if res.Status == 404 {
		res.Errors = append(res.Errors, "User not found")
		return
	} else if res.Status == 400 {
		res.Errors = append(res.Errors, "Error while finding user")
		return
	} else {
		res.Status = 400
	}
	if !auth.CheckPasswordHash(userParams.Password, user.Hash) {
		res.Errors = append(res.Errors, "Invalid password")
		res.Status = 403
		return
	}

	token, success := tokens.EncodeToken(userParams.Username)

	if !success {
		res.Errors = append(res.Errors, "Failed to create token")
		return
	}
	res.Status = 200
	res.Success = true
	res.Response = token
}
