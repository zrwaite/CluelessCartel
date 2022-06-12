import { game, getPx } from "../index.js";
import Button from "./Buttons/Button.js";
import Component from "./Component.js";
import TextSection from "./TextSection.js";
import {Structure} from "../types/structure"
import { LandMaterial } from "../types/hexagon.js";

export class StructureSelector extends Component {
	constructor(parentElement: HTMLElement,	hexagonMaterial: LandMaterial, rowSize: number, setSection: (sectionName:string) => void) {
		super('div', parentElement)
		let rows = 0;
		let allStructures = game.data.AllStructures.filter(structure => {
			if (structure.Name === "Empty" || structure.Natural) return false
			if (structure.LandMaterials.length ===0) return true
			for (let i=0; i<structure.LandMaterials.length; i++) {
				if (structure.LandMaterials[i].Name === hexagonMaterial.Name) {
					return true
				}
			}
			return false
		})
		allStructures.forEach((structure, i) => {
			if (i && i%rowSize === 0) rows++
			let structureButton = new Button(this.element, () => {
				setSection(structure.Name)
			}, {
				position: 'absolute',
				top: getPx(rows*50),
				left: getPx((i%rowSize)*75),
				height: getPx(40), 
				width: getPx(70),
				backgroundColor: 'white'
			})
			let structureName = new TextSection(structureButton.element, 15, structure.Name)
		})
	}
}