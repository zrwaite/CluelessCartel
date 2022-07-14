package user

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/auth/tokens"
	"clueless-cartel-server/database"
	"clueless-cartel-server/database/data"
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
		data.ValidateData(userParams, res)
		if UsernameUsed(userParams.Username) {
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
		token, success := tokens.EncodeToken(user.Username)

		if !success {
			res.Errors = append(res.Errors, "Failed to create token")
			return
		}
		res.Success = true
		res.Status = 201
		user.Hash = ""
		res.Response = struct {
			User  models.PostUser
			Token string
		}{user, token}
	} else {
		userParams.Password = ""
		res.Response = userParams
	}
}
