export interface ResourceCost {
	Resource: Resource
	Cost: number
}

export interface Resource {
	Name: string
}

export interface Resources {
	Metal: number
	Plants: number
	Chemicals: number
}
