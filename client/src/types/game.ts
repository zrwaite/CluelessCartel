import { Drug } from "./drugs";
import {Location} from "./locations";
import { Resource } from "./resources";

export interface GameData {
	AllDrugs:     Drug[]
	AllLocations: Location[]
	AllResources: Resource[]
}