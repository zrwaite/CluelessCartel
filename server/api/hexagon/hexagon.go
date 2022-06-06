package hexagon

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"encoding/json"
	"net/http"
)

func HexagonHandler(r *http.Request, data []byte, res *apiModels.Response) {
	if r.Method != "POST" {
		res.Errors = append(res.Errors, "Method "+r.Method+" is not supported")
		return
	}
	var functionStruct apiModels.FunctionStruct
	devReader := bytes.NewReader(data)
	err := json.NewDecoder(devReader).Decode(&functionStruct)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		switch functionStruct.Function {
		case "buy":
			buyHexagon(data, res)
		default:
			res.Errors = append(res.Errors, "Invalid function")
		}
	}
}
