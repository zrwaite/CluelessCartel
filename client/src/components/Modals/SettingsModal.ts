import { game, getPx, unit } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Button from '../Buttons/Button.js'
import ClickableGridHexagonRow from '../ClickableGridHexagonRow.js'
import Modal from './Modal.js'

export default class SettingsModal extends Modal {
	constructor(hexagonRows: ClickableGridHexagonRow[], styles: StyleObject = {}) {
		super('settings', styles)
		// const hexElement = document.getElementById('hexSection')
		// if (!hexElement) throw Error('hexSection not found')

		const useZoom = () => {
			hexagonRows.forEach((hr) => {
				hr.addStyles({
					height: getPx(hr.height * game.zoom),
					left: getPx(hr.left * game.zoom),
				})
				hr.hexagons.forEach((hex) => {
					let newHeight = getPx(hex.height * game.zoom)
					let newWidth = getPx(hex.width * game.zoom)
					hex.addStyles({
						height: newHeight,
						width: newWidth,
					})
				})
			})
		}
		const zoomIn = () => {
			game.zoom *= 1.1
			useZoom()
		}
		const zoomInButton = new Button(this.element, zoomIn.bind(this), {
			top: getPx(5),
			left: getPx(5),
		})
		zoomInButton.initializeIcon('add.svg')

		const zoomOut = () => {
			game.zoom /= 1.1
			useZoom()
		}
		const zoomOutButton = new Button(this.element, zoomOut.bind(this), {
			top: getPx(45),
			left: getPx(5),
		})
		zoomOutButton.initializeIcon('remove.svg')
		useZoom()
	}
}
