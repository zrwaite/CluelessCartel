export interface Resource {
	Name: string
	SynthesizationCost: ResourcesAmount[]
	SynthesizationTime: number // in minutes
	BatchAmount: number
	StartingUnitPrice: number
}

export interface ResourcesAmount {
	ResourceName: string
	Amount: number
}
