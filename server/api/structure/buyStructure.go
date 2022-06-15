package structure

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/api/base"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"clueless-cartel-server/modules"
	"encoding/json"
)

type BuyStructureParams struct {
	Username      string
	BaseLocation  string
	StructureName string
	HexagonX      int
	HexagonY      int
}

func handleBuyStructure(body []byte, res *apiModels.Response) {
	var buyStructureParams BuyStructureParams
	baseReader := bytes.NewReader(body)
	err := json.NewDecoder(baseReader).Decode(&buyStructureParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		models.ValidateData(buyStructureParams, res)
	}
	if len(res.Errors) != 0 {
		res.Response = buyStructureParams
		return
	}
	var user *models.GetUserReturn
	user, res.Status = dbModules.GetUserGameData(buyStructureParams.Username)
	if res.Status == 404 {
		res.Errors = append(res.Errors, "User not found")
		return
	} else if res.Status == 400 {
		res.Errors = append(res.Errors, "Failed to get user")
		return
	} else {
		res.Status = 400
	}

	location, success := models.GetLocation(buyStructureParams.BaseLocation)
	if !success {
		res.Errors = append(res.Errors, "invalid location")
		return
	}

	// Make sure base exists at location
	base, baseExists := base.GetBase(&user.Bases, location)
	if !baseExists {
		res.Errors = append(res.Errors, "Base doesn't exist at that location")
		return
	}

	structure, success := models.GetStructure(buyStructureParams.StructureName)
	if !success {
		res.Errors = append(res.Errors, "invalid location")
		return
	}

	XIndex, YIndex := modules.GetIndexes(base, res, buyStructureParams.HexagonX, buyStructureParams.HexagonY)
	if len(res.Errors) != 0 {
		return
	}

	hexagon := &base.HexagonRows[YIndex].Hexagons[XIndex]

	/*
		//Find the cost for the new structure
		structureCost := GetStructureCost(base)
		user.Cash += hexagonCost
		if user.Cash < hexagonCost {
			res.Errors = append(res.Errors, "Not enough cash!")
			return
		}

		// Update user cash, and add base
		user.Cash -= hexagonCost
	*/
	if !buyStructure(hexagon, structure, res) {
		return
	}
	success = dbModules.UpdateUser(user)
	if !success {
		res.Errors = append(res.Errors, "Failed to update user - "+err.Error())
		return
	}
	res.Success = true
	res.Status = 200
	res.Response = user
}

func buyStructure(hexagon *models.Hexagon, structure models.Structure, res *apiModels.Response) (success bool) {
	success = false
	// If user already owns hexagon
	if !hexagon.Owned {
		res.Errors = append(res.Errors, "Hexagon not owned")
		return
	}
	// If enemy structures exist, blow em up!
	if hexagon.Structure.Name != "Empty" {
		res.Errors = append(res.Errors, "Hexagon structures must be cleared")
		return
	}

	if structure.Natural {
		res.Errors = append(res.Errors, "Can not buy natural structure")
		return
	}

	if structure.Name == "Empty" {
		res.Errors = append(res.Errors, "Can not buy empty structure")
		return
	}

	validLandMaterial := false
	if len(structure.LandMaterials) == 0 {
		validLandMaterial = true
	} else {
		for _, landMaterial := range structure.LandMaterials {
			if landMaterial.Name == hexagon.LandMaterial.Name {
				validLandMaterial = true
			}
		}
	}
	if !validLandMaterial {
		res.Errors = append(res.Errors, "Structure can't be built on this land material")
		return
	}

	hexagon.Structure = structure

	return true
}
