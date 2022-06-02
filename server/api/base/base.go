package base

import (
	"bytes"
	"clueless-cartel-server/api"
	"clueless-cartel-server/auth"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	api.AddJSONHeader(w)
	if r.URL.Path != "/api/base" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	res := new(api.Response).Init()
	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			res.Errors = append(res.Errors, "Invalid body - "+err.Error())
		} else {
			var functionDevStruct auth.FunctionDevStruct
			devReader := bytes.NewReader(data)
			err := json.NewDecoder(devReader).Decode(&functionDevStruct)
			if err != nil {
				res.Errors = append(res.Errors, "Invalid json - "+err.Error())
			} else {
				if auth.DevCheck(functionDevStruct.Dev, res) {
					switch functionDevStruct.Function {
					case "new":
						newBase(data, res)
					default:
						res.Errors = append(res.Errors, "Invalid function")
					}
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
