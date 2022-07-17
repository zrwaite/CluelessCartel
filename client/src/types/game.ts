
import { LandMaterial } from "./hexagon";
import { Resource } from "./resources";
import { Structure } from "./structure";

export interface GameData {
	AllDrugs:         Resource[]
	AllLocations:     Location[]
	AllResources:     Resource[]
	AllStructures:    Structure[]
	AllLandMaterials: LandMaterial[]
}