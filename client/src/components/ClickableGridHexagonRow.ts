import { getPx } from '../index.js'
import Vector2 from '../modules/Vector2.js'
import ClickableGridHexagon from './ClickableGridHexagon.js'
import Component from './Component.js'
import Modal from './Modals/Modal.js'

export default class ClickableGridHexagonRow extends Component {
	index: number
	height: number = 76
	left: number = 0
	hexagons: ClickableGridHexagon[] = []
	constructor(index: number, parentElement: HTMLElement) {
		super('div', parentElement)
		this.index = index
		this.element.classList.add('clickableHexagonRow')
		this.addStyles({
			position: 'relative',
			height: getPx(76),
		})
		if (index % 2 !== 0) {
			this.addStyles({ left: getPx(50) })
			this.left = 50
		}
	}
	addHexagon(xIndex: number, modal: Modal) {
		this.hexagons.push(new ClickableGridHexagon(new Vector2(xIndex, this.index), this.element, modal, 100))
	}
}
