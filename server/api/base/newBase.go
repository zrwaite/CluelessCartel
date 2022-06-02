package base

import (
	"bytes"
	"clueless-cartel-server/api"
	"clueless-cartel-server/database/models"
	"encoding/json"
)

type PostBaseParams struct {
	Username string
	Location string
}

func newBase(body []byte, res *api.Response) {
	var newBaseParams PostBaseParams
	baseReader := bytes.NewReader(body)
	err := json.NewDecoder(baseReader).Decode(&newBaseParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		models.ValidateData(newBaseParams, res)
	}
	// if len(res.Errors) == 0 {

	// 	_, err := database.MongoDatabase.Collection("users").InsertOne(context.TODO(), user)
	// 	if err != nil {
	// 		res.Errors = append(res.Errors, "Failed to add user to database - "+err.Error())
	// 		return
	// 	}
	// 	res.Success = true
	// 	res.Status = 201
	// 	user.Hash = ""
	// 	res.Response = user
	// } else {
	// 	res.Response = userParams
	// }
	// Get user cash + bases
	// baseCost = 10000 * 10^user.bases.length
	// base.location = body.location
	// base.

}
