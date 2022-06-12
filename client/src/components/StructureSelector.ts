import { game, getPx } from "../index.js";
import Button from "./Buttons/Button.js";
import Component from "./Component.js";
import TextSection from "./TextSection.js";

export class StructureSelector extends Component {
	constructor(parentElement: HTMLElement,	rowSize: number, setSection: (sectionName:string) => void) {
		super('div', parentElement)
		let rows = 0;
		game.data.AllStructures.forEach((structure, i) => {
			if (i && i%rowSize === 0) rows++
			let structureButton = new Button(this.element, () => {
				setSection(structure.Name)
			}, {
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