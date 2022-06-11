import { Drugs } from './drugs'
import { HexagonRow } from './hexagon'
import { Resources } from './resources'
import { Weapons } from './weapons'
import { Location } from './locations'

export interface Base {
	Location: Location
	Index: number
	Resources: Resources
	Drugs: Drugs
	Weapons: Weapons
	HexagonRows: HexagonRow[]
}
