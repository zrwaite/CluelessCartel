import { LandMaterial } from './hexagon'
import { Resources } from './resources'

export interface Structure {
	Moveable: boolean
	Name: string
	LandMaterials: LandMaterial[]
	ResourceCapacity: Resources
}
