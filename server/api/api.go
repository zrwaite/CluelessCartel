package api

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/api/base"
	"clueless-cartel-server/api/gameData"
	"clueless-cartel-server/api/hexagon"
	"clueless-cartel-server/api/signin"
	"clueless-cartel-server/api/structure"
	"clueless-cartel-server/api/user"
	"clueless-cartel-server/auth"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := new(apiModels.Response).Init()
	var data []byte
	if r.Method == "GET" {
		auth.DevCheck(auth.GetDevParam(r), res)
	} else if r.Method == "POST" {
		var err error
		data, err = ioutil.ReadAll(r.Body)
		if err != nil {
			res.Errors = append(res.Errors, "Invalid body - "+err.Error())
		} else {
			var devStruct auth.DevStruct
			devReader := bytes.NewReader(data)
			err := json.NewDecoder(devReader).Decode(&devStruct)
			if err != nil {
				res.Errors = append(res.Errors, "Invalid json - "+err.Error())
			} else {
				auth.DevCheck(devStruct.Dev, res)
			}
		}
	} else if r.Method == "OPTIONS" {
		res.Errors = append(res.Errors, "None")
		res.Status = 200
	} else {
		res.Errors = append(res.Errors, "Method "+r.Method+" is not supported")
		res.Status = 404
	}
	if len(res.Errors) == 0 {
		switch r.URL.Path {
		case "/api/user":
			user.UserHandler(r, data, res)
		case "/api/base":
			base.BaseHandler(r, data, res)
		case "/api/signin":
			signin.SignInHandler(r, data, res)
		case "/api/hexagon":
			hexagon.HexagonHandler(r, data, res)
		case "/api/game":
			gameData.GameDataHandler(r, data, res)
		case "/api/structure":
			structure.StructureHandler(r, data, res)
		default:
			res.Errors = append(res.Errors, "Endpoint not found")
			res.Status = 404
		}
	}
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res)
}
