package api

import (
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/api/base"
	"clueless-cartel-server/api/gameData"
	"clueless-cartel-server/api/hexagon"
	"clueless-cartel-server/api/reset"
	"clueless-cartel-server/api/signin"
	"clueless-cartel-server/api/structure"
	"clueless-cartel-server/api/user"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := new(apiModels.Response).Init()
	var body []byte

	if r.Method == "POST" || r.Method == "DELETE" {
		var err error
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			res.Errors = append(res.Errors, "Invalid body - "+err.Error())
		}
	} else if r.Method == "OPTIONS" {
		res.Errors = append(res.Errors, "None")
		res.Status = 200
	} else if r.Method != "GET" {
		res.Errors = append(res.Errors, "Method "+r.Method+" is not supported")
		res.Status = 404
	}
	if len(res.Errors) == 0 {
		switch r.URL.Path {
		case "/api/user":
			user.UserHandler(r, body, res)
		case "/api/base":
			base.BaseHandler(r, body, res)
		case "/api/signin":
			signin.SignInHandler(r, body, res)
		case "/api/hexagon":
			hexagon.HexagonHandler(r, body, res)
		case "/api/game":
			gameData.GameDataHandler(r, body, res)
		case "/api/structure":
			structure.StructureHandler(r, body, res)
		case "/api/reset":
			reset.ResetHandler(r, body, res)
		default:
			res.Errors = append(res.Errors, "Endpoint not found")
			res.Status = 404
		}
	}
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res)
}
