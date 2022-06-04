package gameData

import (
	"clueless-cartel-server/api/apiModels"
	"net/http"
)

func GameDataHandler(r *http.Request, data []byte, res *apiModels.Response) {
	if r.Method != "GET" {
		res.Errors = append(res.Errors, "Invalid Method")
		return
	}
}

func GetGameData() {

}
