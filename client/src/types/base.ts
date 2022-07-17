import { HexagonRow } from './hexagon'
import { Weapons } from './weapons'
import { Location } from './locations'
import { ResourcesAmount } from './resources'

export interface Base {
	Location: Location
	Index: number
	Resources: ResourcesAmount[]
	Drugs: ResourcesAmount[]
	Weapons: Weapons
	HexagonRows: HexagonRow[]
}