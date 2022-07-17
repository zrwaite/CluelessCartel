export interface Resource {
	Name: string
	SynthesizationCost: ResourcesAmount[]
	SynthesizationTime: number // in minutes
	BatchAmount: number
	StartingUnitPrice: number
	Type: ResourceType
}

export interface ResourceType {
	Name: string
}

export interface ResourcesAmount {
	ResourceName: string
	Amount: number
}
