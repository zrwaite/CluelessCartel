package hexagon

import (
	"clueless-cartel-server/database/models"
	"fmt"
	"math"
)

func GetNumHexagons(hexagonRows *[]models.HexagonRow) (num int) {
	for _, hexagonRow := range *hexagonRows {
		for _, hexagon := range hexagonRow.Hexagons {
			if hexagon.Owned {
				num++
			}
		}
	}
	return
}

func GetHexagonCost(base *models.Base) (num int) {
	// 1000 * 10^index * 1.5^len(hexagons)
	return int(1000 * math.Pow(10, float64(base.Index)) * math.Pow(1.5, float64(GetNumHexagons(&base.HexagonRows)-7)))
}

func ReorganizeHexagons(base *models.Base, centreX int, centreY int) (success bool) {
	var newX int
	var newY int

	// Hexagon Above:
	newX = centreX
	newY = centreY - 1
	if !AddBuyableHexagon(base, newX, newY) {
		return false
	}

	//Hexagon left:
	newX = centreX - 1
	newY = centreY
	if !AddBuyableHexagon(base, newX, newY) {
		return false
	}

	//Hexagon right:
	newX = centreX + 1
	newY = centreY
	if !AddBuyableHexagon(base, newX, newY) {
		return false
	}

	//Hexagon down:
	newX = centreX
	newY = centreY + 1
	if !AddBuyableHexagon(base, newX, newY) {
		return false
	}

	//Diagonals:
	if centreY%2 == 0 {
		newX = centreX - 1
	} else {
		newX = centreX + 1
	}

	//Top diagonal
	newY = centreY + 1
	if !AddBuyableHexagon(base, newX, newY) {
		return false
	}
	//Bottom diagonal
	newY = centreY - 1
	if !AddBuyableHexagon(base, newX, newY) {
		return false
	}
	return true
}

func AddBuyableHexagon(base *models.Base, x int, y int) (success bool) {
	yIndex := -base.HexagonRows[0].Y + y
	xIndex := -base.HexagonRows[0].Hexagons[0].X + x
	if yIndex < 0 {
		AddEmptyHexagonRow(base, true)
		yIndex++
	} else if yIndex >= len(base.HexagonRows) {
		AddEmptyHexagonRow(base, false)
	}
	if xIndex < 0 {
		AddEmptyHexagonColumn(base, true)
		xIndex++
	} else if xIndex >= len(base.HexagonRows[0].Hexagons) {
		AddEmptyHexagonColumn(base, false)
	}
	if yIndex < 0 || xIndex < 0 || yIndex >= len(base.HexagonRows) || xIndex >= len(base.HexagonRows[0].Hexagons) {
		fmt.Println("Invalid indexing while adding buyable hexagon")
		return false
	}
	base.HexagonRows[yIndex].Hexagons[xIndex].Buyable = true
	return true
}

func AddEmptyHexagonRow(base *models.Base, start bool) {
	newLength := len(base.HexagonRows[0].Hexagons)
	startX := base.HexagonRows[0].Hexagons[0].X
	var newY int
	if start {
		newY = base.HexagonRows[0].Y - 1
	} else {
		newY = base.HexagonRows[len(base.HexagonRows)-1].Y + 1
	}
	newHexagonRow := models.HexagonRow{Y: newY}
	for i := startX; i < newLength+startX; i++ {
		newHexagonRow.Hexagons = append(newHexagonRow.Hexagons, models.Hexagon{
			LandMaterial: models.GetLandMaterial(base.Location),
			Structure:    models.EmptyStructure,
			X:            i,
			Owned:        false,
			Buyable:      false,
		})
	}
	if start {
		base.HexagonRows = append([]models.HexagonRow{newHexagonRow}, base.HexagonRows...)
	} else {
		base.HexagonRows = append(base.HexagonRows, newHexagonRow)
	}
}

func AddEmptyHexagonColumn(base *models.Base, start bool) {
	newLength := len(base.HexagonRows)
	startY := base.HexagonRows[0].Y
	var newX int
	if start {
		newX = base.HexagonRows[0].Hexagons[0].X - 1
	} else {
		newX = base.HexagonRows[0].Hexagons[len(base.HexagonRows[0].Hexagons)-1].X + 1
	}
	for i := startY; i < newLength+startY; i++ {
		newHexagon := models.Hexagon{
			LandMaterial: models.GetLandMaterial(base.Location),
			Structure:    models.EmptyStructure,
			X:            newX,
			Owned:        false,
			Buyable:      false,
		}
		if start {
			base.HexagonRows[i-startY].Hexagons = append([]models.Hexagon{newHexagon}, base.HexagonRows[i-startY].Hexagons...)
		} else {
			base.HexagonRows[i-startY].Hexagons = append(base.HexagonRows[i-startY].Hexagons, newHexagon)
		}
	}
}
