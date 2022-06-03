package user

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/database"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"context"
	"encoding/json"
)

func postUser(body []byte, res *apiModels.Response) {
	var userParams models.PostUserParams
	userReader := bytes.NewReader(body)
	err := json.NewDecoder(userReader).Decode(&userParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		models.ValidateData(userParams, res)
		if usernameUsed(userParams.Username) {
			res.Errors = append(res.Errors, "Username in use")
		}
	}
	if len(res.Errors) == 0 {
		var user models.PostUser
		user.InitData(&userParams)
		err = user.CreateHash(userParams.Password)
		if err != nil {
			res.Errors = append(res.Errors, "Failed to hash password - "+err.Error())
			return
		}
		_, err := database.MongoDatabase.Collection("users").InsertOne(context.TODO(), user)
		if err != nil {
			res.Errors = append(res.Errors, "Failed to add user to database - "+err.Error())
			return
		}
		res.Success = true
		res.Status = 201
		user.Hash = ""
		res.Response = user
	} else {
		userParams.Password = ""
		res.Response = userParams
	}
}

func usernameUsed(username string) bool {
	_, status := dbModules.GetUserData(username, models.GetIdOpts)
	if status == 404 {
		return false
	} else if status == 200 {
		return true
	}
	return false
}
