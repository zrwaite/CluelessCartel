package user

import (
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/database/dbModules"
	"net/http"
)

func UserHandler(r *http.Request, body []byte, res *apiModels.Response) {
	if r.Method == "POST" {
		postUser(body, res)
	} else {
		getUser(r, res)
	}
}

func UsernameUsed(username string) bool {
	_, status := dbModules.GetUserIdData(username)
	if status == 404 {
		return false
	} else if status == 200 {
		return true
	}
	return false
}
