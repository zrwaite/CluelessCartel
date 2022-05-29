package user

import (
	"clueless-cartel-server/api"
	"clueless-cartel-server/auth"
	"clueless-cartel-server/database"
	"clueless-cartel-server/database/models"
	"context"
	"encoding/json"
	"net/http"
)

func postUser(r *http.Request, res *api.Response) {
	var userParams models.PostUserParams
	err := json.NewDecoder(r.Body).Decode(&userParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
		return
	} else {
		auth.DevCheck(userParams.Dev, res)
		models.ValidateData(userParams, res)
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
		res.Response = user
	} else {
		res.Response = userParams
	}
}
