import { ResourceCost } from './resources'

export interface Drug {
	Cost: ResourceCost[]
	BatchAmount: number
	UnitPrice: number
}

export interface Drugs {
	Opioids: number
	Weed: number
	Meth: number
}
