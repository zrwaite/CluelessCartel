package gameData

import (
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/database/data"
	"net/http"
)

func GameDataHandler(r *http.Request, body []byte, res *apiModels.Response) {
	if r.Method != "GET" {
		res.Errors = append(res.Errors, "Invalid Method")
		return
	}
	res.Response = data.GetGameData()
	res.Status = 200
	res.Success = true
}
