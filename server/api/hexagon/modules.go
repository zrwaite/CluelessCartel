package hexagon

import (
	"clueless-cartel-server/database/data"
	"clueless-cartel-server/database/models"
	"fmt"
	"math"
	"math/rand"
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
	hexagon := &base.HexagonRows[yIndex].Hexagons[xIndex]

	if !hexagon.Buyable {
		hexagon.Buyable = true
		surroundingHexagons := GetSurroundingHexagons(base, x, y)
		enemyCamp := false
		forest := false
		for _, hex := range surroundingHexagons {
			if hex.Structure.Enemy {
				enemyCamp = true
			}
			if hex.Structure.Name == "Trees" {
				forest = true
			}
		}
		landMaterial := data.GetSemiRandomLandMaterial(base.Location, surroundingHexagons)

		hexagon.Structure = data.GetSemiRandomStructure(landMaterial, enemyCamp, forest)
		hexagon.LandMaterial = landMaterial
	}

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
			Structure:    data.EmptyStructure,
			LandMaterial: data.DeadLand,
			Rotation:     rand.Intn(6),
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
			Structure:    data.EmptyStructure,
			LandMaterial: data.DeadLand,
			Rotation:     rand.Intn(6),
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

func GetSurroundingHexagons(base *models.Base, x int, y int) (hexagons []models.Hexagon) {
	yIndex := -base.HexagonRows[0].Y + y
	xIndex := -base.HexagonRows[0].Hexagons[0].X + x
	var newX int
	var newY int
	// Hexagon Above:
	newX = xIndex
	newY = yIndex - 1
	if ValidHexIndex(base, newX, newY) {
		newHex := base.HexagonRows[newY].Hexagons[newX]
		if newHex.Buyable {
			hexagons = append(hexagons, newHex)
		}
	}

	//Hexagon left:
	newX = xIndex - 1
	newY = yIndex
	if ValidHexIndex(base, newX, newY) {
		newHex := base.HexagonRows[newY].Hexagons[newX]
		if newHex.Buyable {
			hexagons = append(hexagons, newHex)
		}
	}

	//Hexagon right:
	newX = xIndex + 1
	newY = yIndex
	if ValidHexIndex(base, newX, newY) {
		newHex := base.HexagonRows[newY].Hexagons[newX]
		if newHex.Buyable {
			hexagons = append(hexagons, newHex)
		}
	}

	//Hexagon down:
	newX = xIndex
	newY = yIndex + 1
	if ValidHexIndex(base, newX, newY) {
		newHex := base.HexagonRows[newY].Hexagons[newX]
		if newHex.Buyable {
			hexagons = append(hexagons, newHex)
		}
	}

	//Diagonals:
	if y%2 == 0 {
		newX = xIndex - 1
	} else {
		newX = xIndex + 1
	}

	//Top diagonal
	newY = yIndex + 1
	if ValidHexIndex(base, newX, newY) {
		newHex := base.HexagonRows[newY].Hexagons[newX]
		if newHex.Buyable {
			hexagons = append(hexagons, newHex)
		}
	}
	//Bottom diagonal
	newY = yIndex - 1
	if ValidHexIndex(base, newX, newY) {
		newHex := base.HexagonRows[newY].Hexagons[newX]
		if newHex.Buyable {
			hexagons = append(hexagons, newHex)
		}
	}
	return
}

func ValidHexIndex(base *models.Base, xIndex int, yIndex int) (valid bool) {
	return !(xIndex < 0 || yIndex < 0 || yIndex >= len(base.HexagonRows) || xIndex >= len(base.HexagonRows[yIndex].Hexagons))
}
