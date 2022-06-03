package user

import (
	"clueless-cartel-server/api/apiModels"
	"net/http"
)

func UserHandler(r *http.Request, data []byte, res *apiModels.Response) {
	if r.Method == "POST" {
		postUser(data, res)
	} else {
		getUser(r, res)
	}
}
