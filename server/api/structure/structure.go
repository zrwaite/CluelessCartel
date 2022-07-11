package structure

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"encoding/json"
	"net/http"
)

func StructureHandler(r *http.Request, data []byte, res *apiModels.Response) {
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
			handleBuyStructure(data, res)
		case "remove":
			handleRemoveStructure(data, res)
		case "explode":
			handleExplodeStructure(data, res)
		case "driveby":
			handleDriveByStructure(data, res)
		default:
			res.Errors = append(res.Errors, "Invalid function")
		}
	}
}
