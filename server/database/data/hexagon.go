package data

import (
	"clueless-cartel-server/database/models"
	"math/rand"
)

var DeadLand = models.LandMaterial{Name: "Dead"}
var Grass = models.LandMaterial{"Grass"}
var Dirt = models.LandMaterial{"Dirt"}
var Sand = models.LandMaterial{"Sand"}
var Pavement = models.LandMaterial{"Pavement"}
var Water = models.LandMaterial{"Water"}
var AllLandMaterials = []models.LandMaterial{Grass, Dirt, Sand, Pavement, Water}

func ValidateHexagonRowsStructure(hexagonRows *[]models.HexagonRow) (valid bool) {
	if len(*hexagonRows) < 5 {
		return false
	}
	length := len((*hexagonRows)[0].Hexagons)
	if length < 5 {
		return false
	}
	for _, hexagonRow := range *hexagonRows {
		newLength := len(hexagonRow.Hexagons)
		if length != newLength {
			return false
		}
	}
	return true
}

func CreateStartingHexagonRows(location models.Location) []models.HexagonRow {
	var hexagonRows = []models.HexagonRow{}
	for i1 := -2; i1 <= 2; i1++ {
		/* Create hexagon Grid:
		  ❌ ⬜ ⬜ ⬜ ❌
		   ⬜ ⬛ ⬛ ⬜ ❌
		 ⬜ ⬛ ⬛ ⬛ ⬜
		  ⬜ ⬛ ⬛ ⬜ ❌
		❌ ⬜ ⬜ ⬜ ❌
		*/
		var hexagonRow = models.HexagonRow{Y: i1}
		for i2 := -2; i2 <= 2; i2++ {
			owned := false
			buyable := true
			if i1 == -2 || i1 == 2 {
				if i2 == -2 || i2 == 2 {
					buyable = false
				}
			} else if i1 == -1 || i1 == 1 {
				if i2 == 2 {
					buyable = false
				} else if i2 == -1 || i2 == 0 {
					owned = true
				}
			} else if i2 == -1 || i2 == 0 || i2 == 1 {
				owned = true
			}
			hexagonRow.Hexagons = append(hexagonRow.Hexagons, models.Hexagon{
				LandMaterial: location.LandMaterial,
				Structure:    EmptyStructure,
				Owned:        owned,
				Rotation:     rand.Intn(6),
				Buyable:      buyable,
				X:            i2,
			})
		}
		hexagonRows = append(hexagonRows, hexagonRow)
	}
	return hexagonRows
}
