import { Resources } from './resources'

export interface Structure {
	Moveable: boolean
	Name: string
	Image: string
	LandMaterials: string[]
	ResourceCapacity: Resources
}
