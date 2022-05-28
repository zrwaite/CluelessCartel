package auth

import (
	"clueless-cartel-server/api"
	"clueless-cartel-server/settings"
	"encoding/json"
	"net/http"
)

func DevCheck(r *http.Request, res *api.Response) bool {
	var devQuery = struct {
		Dev string
	}{}
	if r.URL.Path == "GET" {
		devQuery.Dev = r.URL.Query().Get("dev")
		if devQuery.Dev == "" {
			res.Errors = append(res.Errors, "dev query param not defined")
			return false
		}
	} else {
		err := json.NewDecoder(r.Body).Decode(&devQuery)
		if err != nil || devQuery.Dev == "" {
			res.Errors = append(res.Errors, "dev param not defined")
			return false
		}
	}
	if devQuery.Dev != "true" && settings.DEV {
		res.Errors = append(res.Errors, "Frontend is in dev mode but backend is in production mode")
		return false
	} else if devQuery.Dev != "false" && !settings.DEV {
		res.Errors = append(res.Errors, "Frontend is in production mode but backend is in dev mode")
		return false
	}
	return true
}
