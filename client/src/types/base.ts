import { Drugs } from './drugs'
import { HexagonRow } from './hexagon'
import { Resources } from './resources'
import { Weapons } from './weapons'

export type Location = 'Nevada' | 'Florida' | 'Canadian Border' | 'New York'

export interface Base {
	Location: string
	Index: number
	Resources: Resources
	Drugs: Drugs
	Weapons: Weapons
	HexagonRows: HexagonRow[]
}
