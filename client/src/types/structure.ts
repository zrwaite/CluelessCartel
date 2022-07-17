import { LandMaterial } from './hexagon'
import { ResourcesAmount } from './resources'

export interface Structure {
	Moveable: boolean
	Name: string
	Enemy: boolean,
	Natural: boolean,
	LandMaterials: LandMaterial[]
	ResourceCapacity: ResourcesAmount[]
}
