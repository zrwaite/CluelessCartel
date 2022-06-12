import { Drug } from "./drugs";
import {Location} from "./locations";
import { Resource } from "./resources";
import { Structure } from "./structure";

export interface GameData {
	AllDrugs:     Drug[]
	AllLocations: Location[]
	AllResources: Resource[]
	AllStructures: Structure[]
}