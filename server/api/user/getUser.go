package user

import (
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"

	"net/http"
)

func getUser(r *http.Request, res *apiModels.Response) {
	username := r.URL.Query().Get("username")
	if username == "" {
		res.Errors = append(res.Errors, "Username not defined")
		return
	}
	res.Response, res.Status = dbModules.GetUserData(username, models.GetUserOpts)
	if res.Status == 404 {
		res.Errors = append(res.Errors, "User not found")
	} else if res.Status == 200 {
		res.Success = true
	} else {
		res.Errors = append(res.Errors, "Failed to get user")
	}
}
