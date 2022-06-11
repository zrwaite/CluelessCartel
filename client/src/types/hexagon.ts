import { Structure } from './structure'


export interface LandMaterial {
	Name: string
}

export interface Hexagon {
	LandMaterial: LandMaterial
	Structure: Structure
	X: number
	Y?: number
	Owned: boolean
	Buyable: boolean
}

export interface HexagonRow {
	Hexagons: Hexagon[]
	Y: number
}
