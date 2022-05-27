package user

import (
	"clueless-cartel-server/api"
	"clueless-cartel-server/auth"
	"clueless-cartel-server/database"

	"context"
	"encoding/json"
	"net/http"
)

type PostUserParams struct {
	Username string
	Password string
}

type PostUser struct {
	Username string
	Hash     string
}

func (user *PostUser) parseData(userParams *PostUserParams) {
	user.Username = userParams.Username
}

func (user *PostUser) createHash(password string) error {
	var err error
	user.Hash, err = auth.HashPassword(password)
	return err
}

func postUser(r *http.Request, res *api.Response) {
	var userParams PostUserParams
	err := json.NewDecoder(r.Body).Decode(&userParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
		return
	} else {
		if userParams.Username == "" {
			res.Errors = append(res.Errors, "Missing username")
		}
		if userParams.Password == "" {
			res.Errors = append(res.Errors, "Missing password")
		}
	}
	if len(res.Errors) == 0 {
		var user PostUser
		user.parseData(&userParams)
		err = user.createHash(userParams.Password)
		if err != nil {
			res.Errors = append(res.Errors, "Failed to hash password - "+err.Error())
			return
		}
		usersCollection := database.MongoDatabase.Collection("users")
		if err != nil {
			res.Errors = append(res.Errors, "Failed to convert user to BSON - "+err.Error())
			return
		}
		_, err := usersCollection.InsertOne(context.TODO(), user)
		if err != nil {
			res.Errors = append(res.Errors, "Failed to add user to database - "+err.Error())
			return
		}
		res.Success = true
		res.Status = 201
	}
	res.Response = userParams
}
