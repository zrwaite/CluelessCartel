import { game, getPx, unit } from '../../index.js'
import { gameAPI } from '../../modules/gameAPI.js'
import { Hexagon } from '../../types/hexagon.js'
import { StyleObject } from '../../types/styles.js'
import Button from '../Buttons/Button.js'
import Modal from './Modal.js'

export default class HexagonModal extends Modal {
	zoom: number = 1
	buyButton?: Button
	hexagon: Hexagon
	constructor(hexagon: Hexagon, styles: StyleObject = {}) {
		super(`${hexagon.X}-${hexagon.Y}`, styles)
		this.hexagon = hexagon
		if (hexagon.Buyable && !hexagon.Owned) {
			this.buyButton = new Button(this.element, this.buyHexagon.bind(this))
			this.buyButton.element.innerText = `Buy ${hexagon.X}, ${hexagon.Y}`
		}
		let text = document.createElement('p')
		text.innerText = `${hexagon.X}, ${hexagon.Y}`
		this.element.appendChild(text)
	}
	async buyHexagon() {
		const response = await gameAPI('/hexagon', 'POST', {
			username: "Insomnizac5",
			function: "buy",
			BaseLocation: "New York",
			HexagonX: this.hexagon.X,
			HexagonY: this.hexagon.Y,
		})
		if (response) {
			if (response.Success) {
				game.data = response.Response
			} else {
				alert(JSON.stringify(response.Errors))
			}
		} else throw Error("Failed to get user!")
		game.start()
	}
}
