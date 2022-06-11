import { LandMaterial } from './hexagon'
import { Resources } from './resources'

export interface Structure {
	Moveable: boolean
	Name: string
	Image: string
	LandMaterials: LandMaterial[]
	ResourceCapacity: Resources
}
