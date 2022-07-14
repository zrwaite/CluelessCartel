package structure

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"encoding/json"
	"net/http"
)

func StructureHandler(r *http.Request, body []byte, res *apiModels.Response) {
	if r.Method != "POST" {
		res.Errors = append(res.Errors, "Method "+r.Method+" is not supported")
		return
	}
	var functionStruct apiModels.FunctionStruct
	devReader := bytes.NewReader(body)
	err := json.NewDecoder(devReader).Decode(&functionStruct)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		switch functionStruct.Function {
		case "buy":
			handleBuyStructure(body, res)
		case "remove":
			handleRemoveStructure(body, res)
		case "explode":
			handleExplodeStructure(body, res)
		case "driveby":
			handleDriveByStructure(body, res)
		default:
			res.Errors = append(res.Errors, "Invalid function")
		}
	}
}
