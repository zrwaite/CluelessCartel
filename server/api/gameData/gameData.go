package gameData

import (
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/database/models"
	"net/http"
)

func GameDataHandler(r *http.Request, data []byte, res *apiModels.Response) {
	if r.Method != "GET" {
		res.Errors = append(res.Errors, "Invalid Method")
		return
	}
	res.Response = models.GetGameData()
	res.Status = 200
	res.Success = true
}
