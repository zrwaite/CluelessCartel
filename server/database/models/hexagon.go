package models

var LandMaterial = struct {
	Dirt     int
	Grass    int
	Sand     int
	Pavement int
}{0, 1, 2, 3}

type Hexagon struct {
	LandMaterial int
	Structure    Structure
	Index        int
}

type HexagonRow struct {
	index    int
	Hexagons []Hexagon
}

var StartingHexagonRows = [3]HexagonRow{
	{
		index: -1,
		Hexagons: []Hexagon{
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        -1,
				Structure:    EmptyStructure,
			},
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        0,
				Structure:    EmptyStructure,
			},
		},
	},
	{
		index: 0,
		Hexagons: []Hexagon{
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        -1,
				Structure:    EmptyStructure,
			},
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        0,
				Structure:    RV,
			},
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        1,
				Structure:    EmptyStructure,
			},
		},
	},
	{
		index: 1,
		Hexagons: []Hexagon{
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        -1,
				Structure:    EmptyStructure,
			},
			{
				LandMaterial: LandMaterial.Dirt,
				Index:        0,
				Structure:    EmptyStructure,
			},
		},
	},
}
