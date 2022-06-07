import { getPx, unit } from '../../index.js'
import { Hexagon } from '../../types/hexagon.js'
import { StyleObject } from '../../types/styles.js'
import Button from '../Buttons/Button.js'
import Modal from './Modal.js'

export default class HexagonModal extends Modal {
	zoom: number = 1
	buyButton?: Button
	constructor(hexagon: Hexagon, styles: StyleObject = {}) {
		super(styles)
		if (hexagon.Buyable && !hexagon.Owned) {
			this.buyButton = new Button(this.element, () => alert('Buy'))
			this.buyButton.element.innerText = `Buy ${hexagon.X}, ${hexagon.Y}`
		}
	}
}
