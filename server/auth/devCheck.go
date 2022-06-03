package auth

import (
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/settings"
	"net/http"
)

type DevStruct struct {
	Dev string
}

func DevCheck(frontendDev string, res *apiModels.Response) bool {
	if frontendDev == "" {
		res.Errors = append(res.Errors, "dev parameter not defined")
	}
	if frontendDev != "true" && settings.DEV {
		res.Errors = append(res.Errors, "Frontend is in dev mode but backend is in production mode")
		return false
	} else if frontendDev != "false" && !settings.DEV {
		res.Errors = append(res.Errors, "Frontend is in production mode but backend is in dev mode")
		return false
	}
	return true
}

func GetDevParam(r *http.Request) string {
	return r.URL.Query().Get("dev")
}
