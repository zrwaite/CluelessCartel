package hexagon

import (
	"bytes"
	"clueless-cartel-server/api/apiModels"
	"clueless-cartel-server/api/base"
	"clueless-cartel-server/database/dbModules"
	"clueless-cartel-server/database/models"
	"encoding/json"
)

type BuyHexagonParams struct {
	Username     string
	BaseLocation string
	HexagonX     int
	HexagonY     int
}

func buyHexagon(body []byte, res *apiModels.Response) {
	var buyHexagonParams BuyHexagonParams
	baseReader := bytes.NewReader(body)
	err := json.NewDecoder(baseReader).Decode(&buyHexagonParams)
	if err != nil {
		res.Errors = append(res.Errors, "Invalid json - "+err.Error())
	} else {
		models.ValidateData(buyHexagonParams, res)
	}
	if len(res.Errors) == 0 {
		var user *models.GetUserReturn
		user, res.Status = dbModules.GetUserGameData(buyHexagonParams.Username)
		if res.Status == 404 {
			res.Errors = append(res.Errors, "User not found")
			return
		} else if res.Status == 400 {
			res.Errors = append(res.Errors, "Failed to get user")
			return
		} else {
			res.Status = 400
		}

		location, success := models.GetLocation(buyHexagonParams.BaseLocation)
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

		/*
			//Find the cost for the new hexagon
			hexagonCost := GetHexagonCost(base)
			if user.Cash < hexagonCost {
				res.Errors = append(res.Errors, "Not enough cash!")
				return
			}

			// Update user cash, and add base
			user.Cash -= hexagonCost
		*/

		// Attempt to buy hexagon, update user on success
		if !BuyHexagon(base, buyHexagonParams.HexagonX, buyHexagonParams.HexagonY, res) {
			return
		}
		if !models.ValidateHexagonRowsStructure(&base.HexagonRows) {
			res.Errors = append(res.Errors, "Failed to create base")
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
	} else {
		res.Response = buyHexagonParams
	}
}

func BuyHexagon(base *models.Base, HexagonX int, HexagonY int, res *apiModels.Response) (success bool) {
	if !models.ValidateHexagonRowsStructure(&base.HexagonRows) {
		res.Errors = append(res.Errors, "Base is in corrupted state")
		return
	}
	YIndex := -base.HexagonRows[0].Y + HexagonY
	XIndex := -base.HexagonRows[0].Hexagons[0].X + HexagonX

	// Validate range
	if YIndex < 0 || YIndex >= len(base.HexagonRows) {
		res.Errors = append(res.Errors, "Invalid YIndex")
		return
	}
	if XIndex < 0 || XIndex >= len(base.HexagonRows[0].Hexagons) {
		res.Errors = append(res.Errors, "Invalid XIndex")
		return
	}

	buyingHexagon := &base.HexagonRows[YIndex].Hexagons[XIndex]

	// If user already owns hexagon
	if buyingHexagon.Owned {
		res.Errors = append(res.Errors, "Hexagon already owned")
		return
	}

	// If hexagon is not purchaseable
	if !buyingHexagon.Buyable {
		res.Errors = append(res.Errors, "Hexagon not within purchase range")
		return
	}

	// If enemy structures exist, blow em up!
	if buyingHexagon.Structure.Enemy {
		res.Errors = append(res.Errors, "Enemy hexagon structures must be cleared")
		return
	}

	if !ReorganizeHexagons(base, HexagonX, HexagonY) {
		res.Errors = append(res.Errors, "Failed to restructure hexagons")
		return false
	}
	YIndex = -base.HexagonRows[0].Y + HexagonY
	XIndex = -base.HexagonRows[0].Hexagons[0].X + HexagonX
	base.HexagonRows[YIndex].Hexagons[XIndex].Owned = true

	return true
}
