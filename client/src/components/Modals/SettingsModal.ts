import { getPx, unit } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Button from '../Buttons/Button.js'
import ClickableGridHexagonRow from '../ClickableGridHexagonRow.js'
import Modal from './Modal.js'

export default class SettingsModal extends Modal {
	zoom: number = 1
	constructor(parentElement: HTMLElement, hexagonRows: ClickableGridHexagonRow[], styles: StyleObject = {}) {
		super(parentElement, styles)
		// const hexElement = document.getElementById('hexSection')
		// if (!hexElement) throw Error('hexSection not found')

		const zoomIn = () => {
			this.zoom *= 1.1
			hexagonRows.forEach((hr) => {
				hr.addStyles({
					height: getPx(hr.height * this.zoom),
					left: getPx(hr.left * this.zoom),
				})
				hr.hexagons.forEach((hex) => {
					let newSize = getPx(hex.size * this.zoom)
					hex.addStyles({
						backgroundSize: newSize + ' ' + newSize,
						height: newSize,
						width: newSize,
					})
				})
			})
		}
		const zoomInButton = new Button(this.element, zoomIn.bind(this), {
			top: getPx(5),
			left: getPx(5),
		})
		zoomInButton.initializeIcon('add.svg')

		const zoomOut = () => {
			this.zoom /= 1.1
			hexagonRows.forEach((hr) => {
				hr.addStyles({
					height: getPx(hr.height * this.zoom),
					left: getPx(hr.left * this.zoom),
				})
				hr.hexagons.forEach((hex) => {
					let newSize = getPx(hex.size * this.zoom)
					hex.addStyles({
						backgroundSize: newSize + ' ' + newSize,
						height: newSize,
						width: newSize,
					})
				})
			})
		}
		const zoomOutButton = new Button(this.element, zoomOut.bind(this), {
			top: getPx(45),
			left: getPx(5),
		})
		zoomOutButton.initializeIcon('remove.svg')
	}
}
