import { game, getPx, unit } from '../../index.js'
import { buyHexagon } from '../../modules/api/buyHexagon.js'
import { buyStructure } from '../../modules/api/buyStructure.js'
import { gameAPI } from '../../modules/gameAPI.js'
import { Hexagon } from '../../types/hexagon.js'
import { StyleObject } from '../../types/styles.js'
import Button from '../Buttons/Button.js'
import { StructureSelector } from '../StructureSelector.js'
import TextSection from '../TextSection.js'
import Modal from './Modal.js'

export default class HexagonModal extends Modal {
	buyButton?: Button
	buyStructureButton?: Button
	buyStructureButtonText?: TextSection
	structureName?: string
	hexagon: Hexagon
	constructor(hexagon: Hexagon, styles: StyleObject = {}) {
		super(`${hexagon.X}-${hexagon.Y}`, styles)
		this.hexagon = hexagon
		if (hexagon.Buyable && !hexagon.Owned) {
			this.buyButton = new Button(this.element, this.buyHexagon.bind(this))
			new TextSection(this.buyButton.element, 15, `Buy ${hexagon.X}, ${hexagon.Y}`)
		}
		if (hexagon.Owned) {
			let structureSection = new StructureSelector(this.element, hexagon.LandMaterial, 5, this.setStructureName.bind(this))
			this.buyStructureButton = new Button(this.element, this.buyStructure.bind(this), {
				top: getPx(100),
				width: getPx(100)
			})
			this.buyStructureButtonText = new TextSection(this.buyStructureButton.element, 15, 'Select Structure')
		}
		new TextSection(this.element, 10, `${hexagon.X}, ${hexagon.Y}`)
	}
	setStructureName(name: string) {
		this.structureName = name
		this.buyStructureButtonText?.setText('Buy ' + name)
	}
	buyHexagon() {
		buyHexagon(this.hexagon.X, this.hexagon.Y)
	}
	buyStructure() {
		if (!this.structureName) {
			alert("structure not selected")
			return
		}
		buyStructure(this.hexagon.X, this.hexagon.Y, this.structureName)	
	}
}
