import { getPx } from '../index.js'
import Vector2 from '../modules/Vector2.js'
import { HexagonRow } from '../types/hexagon.js'
import ClickableGridHexagon from './ClickableGridHexagon.js'
import Component from './Component.js'
import Modal from './Modals/Modal.js'

export default class ClickableGridHexagonRow extends Component {
	index: number
	height: number = 86
	left: number = 0
	hexagons: ClickableGridHexagon[] = []
	hexagonRow: HexagonRow
	constructor(index: number, parentElement: HTMLElement, hexagonRow: HexagonRow) {
		super('div', parentElement)
		this.index = index
		this.hexagonRow = hexagonRow
		this.element.classList.add('clickableHexagonRow')
		this.addStyles({
			position: 'relative',
			height: getPx(this.height),
		})
		if (index % 2 !== 0) {
			this.addStyles({ left: getPx(50) })
			this.left = 50
		}
	}
	initializeHexagons() {
		this.hexagonRow.Hexagons.forEach((hexagon, i) => {
			hexagon.Y = this.hexagonRow.Y
			this.hexagons.push(new ClickableGridHexagon(this.element, hexagon))
		})
	}
}
