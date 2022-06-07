import { Structure } from './structure'

export type LandMaterial = 'Dirt' | 'Sand' | 'Grass' | 'Pavement'

export interface Hexagon {
	LandMaterial: LandMaterial
	Structure: Structure
	X: number
	Owned: boolean
	Buyable: boolean
}

export interface HexagonRow {
	Hexagons: Hexagon[]
	Y: number
}
