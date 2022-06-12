import { LandMaterial } from './hexagon'
import { Resources } from './resources'

export interface Structure {
	Moveable: boolean
	Name: string
	Enemy: boolean,
	Natural: boolean,
	LandMaterials: LandMaterial[]
	ResourceCapacity: Resources
}
