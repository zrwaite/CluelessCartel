package auth

import (
	"clueless-cartel-server/api"
	"clueless-cartel-server/settings"
	"fmt"
	"net/http"
)

type DevStruct struct {
	Dev string
}

type FunctionDevStruct struct {
	Function string
	DevStruct
}

func DevCheck(frontendDev string, res *api.Response) bool {
	fmt.Println(frontendDev)
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
