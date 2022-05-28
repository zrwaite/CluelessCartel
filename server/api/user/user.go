package user

import (
	"clueless-cartel-server/api"
	"clueless-cartel-server/auth"
	"encoding/json"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	api.AddJSONHeader(w)
	if r.URL.Path != "/api/user" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	res := new(api.Response).Init()
	if auth.DevCheck(r, res) {
		if r.Method == "GET" {
			getUser(r, res)
		} else if r.Method == "POST" {
			postUser(r, res)
		} else {
			res.Errors = append(res.Errors, "Method "+r.Method+" is not supported")
			res.Status = 404
		}
	}
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res)
}
