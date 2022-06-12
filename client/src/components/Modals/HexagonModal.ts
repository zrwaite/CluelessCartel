import { game, getPx, unit } from '../../index.js'
import { gameAPI } from '../../modules/gameAPI.js'
import { Hexagon } from '../../types/hexagon.js'
import { StyleObject } from '../../types/styles.js'
import Button from '../Buttons/Button.js'
import TextSection from '../TextSection.js'
import Modal from './Modal.js'

export default class HexagonModal extends Modal {
	buyButton?: Button
	hexagon: Hexagon
	constructor(hexagon: Hexagon, styles: StyleObject = {}) {
		super(`${hexagon.X}-${hexagon.Y}`, styles)
		this.hexagon = hexagon
		if (hexagon.Buyable && !hexagon.Owned) {
			this.buyButton = new Button(this.element, this.buyHexagon.bind(this))
			new TextSection(this.buyButton.element, 15, `Buy ${hexagon.X}, ${hexagon.Y}`)
		}
		new TextSection(this.element, 10, `${hexagon.X}, ${hexagon.Y}`)
	}
	async buyHexagon() {
		const response = await gameAPI('/hexagon', 'POST', {
			username: game.user.Username,
			function: 'buy',
			BaseLocation: game.base?.Location.Name,
			HexagonX: this.hexagon.X,
			HexagonY: this.hexagon.Y,
		})
		if (response) {
			if (response.Success) {
				game.user = response.Response
			} else {
				alert(JSON.stringify(response.Errors))
			}
		} else throw Error('Failed to get user!')
		game.trySaveScroll()
		game.start()
	}
}
