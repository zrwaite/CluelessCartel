import { game, getPx, unit } from '../../index.js'
import { buyHexagon } from '../../modules/api/buyHexagon.js'
import { buyStructure } from '../../modules/api/buyStructure.js'
import { explodeStructure } from '../../modules/api/explodeStructure.js'
import { driveByStructure } from '../../modules/api/driveByStructure.js'
import { removeStructure } from '../../modules/api/removeStructure.js'
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
	driveByStructureButton?: Button
	explodeStructureButton?: Button
	buyStructureButtonText?: TextSection
	structureName?: string
	hexagon: Hexagon
	constructor(hexagon: Hexagon, styles: StyleObject = {}) {
		super(`${hexagon.X}-${hexagon.Y}`, styles)
		this.hexagon = hexagon
		let title = new TextSection(this.element, 20, "Plot of " + hexagon.LandMaterial.Name, {
			width: '100%',
			textAlign: 'center'
		})
		let structureText = new TextSection(this.element, 14, "" + hexagon.Structure.Name,{

		})

		if (hexagon.Buyable && !hexagon.Owned) {
			if (hexagon.Structure.Enemy) {
				title.setText("Ememy " + title.getText())
				this.driveByStructureButton = new Button(this.element, this.driveByStructure.bind(this))
				new TextSection(this.driveByStructureButton.element, 15, `Shoot up ${hexagon.X}, ${hexagon.Y}`)
				
				this.explodeStructureButton = new Button(this.element, this.explodeStructure.bind(this), {
					left: getPx(100)
				})

				new TextSection(this.explodeStructureButton.element, 15, `Explode ${hexagon.X}, ${hexagon.Y}`)

			} else {
				this.buyButton = new Button(this.element, this.buyHexagon.bind(this))
				new TextSection(this.buyButton.element, 15, `Buy ${hexagon.X}, ${hexagon.Y}`)
			}
		}
		if (hexagon.Owned ) {
			if (hexagon.Structure.Name === "Empty") {
				//Buy menu
				let structureSection = new StructureSelector(this.element, hexagon.LandMaterial, 5, this.setStructureName.bind(this), {
					position: 'absolute',
					top: getPx(50)	
				})
				this.buyStructureButton = new Button(this.element, this.buyStructure.bind(this), {
					top: getPx(100),
					width: getPx(100)
				})
				this.buyStructureButtonText = new TextSection(this.buyStructureButton.element, 15, 'Select Structure')
			} else {
				let removeStructureButton = new Button(this.element, this.removeStructure.bind(this), {
					top: getPx(150),
					width: getPx(100)
				})
				let removeStrucureButtonText = new TextSection(removeStructureButton.element, 15, 'Remove Structure')
			}
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
	explodeStructure() {
		explodeStructure(this.hexagon.X, this.hexagon.Y)
	}
	driveByStructure() {
		driveByStructure(this.hexagon.X, this.hexagon.Y)
	}
	buyStructure() {
		if (!this.structureName) {
			alert("structure not selected")
			return
		}
		buyStructure(this.hexagon.X, this.hexagon.Y, this.structureName)	
	}
	removeStructure() {
		removeStructure(this.hexagon.X, this.hexagon.Y)
	}
}
