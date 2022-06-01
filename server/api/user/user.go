package user

import (
	"bytes"
	"clueless-cartel-server/api"
	"clueless-cartel-server/auth"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	api.AddJSONHeader(w)
	if r.URL.Path != "/api/user" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	res := new(api.Response).Init()
	if r.Method == "GET" {
		if auth.DevCheck(auth.GetDevParam(r), res) {
			getUser(r, res)
		}
	} else if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			res.Errors = append(res.Errors, "Invalid body - "+err.Error())
		} else {
			var devStruct auth.DevStruct
			devReader := bytes.NewReader(data)
			err := json.NewDecoder(devReader).Decode(&devStruct)
			if err != nil {
				res.Errors = append(res.Errors, "Invalid json - "+err.Error())
			} else {
				if auth.DevCheck(devStruct.Dev, res) {
					postUser(data, res)
				}
			}
		}
	} else {
		res.Errors = append(res.Errors, "Method "+r.Method+" is not supported")
		res.Status = 404
	}
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res)
}
