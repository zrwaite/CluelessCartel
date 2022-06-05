import { getPx } from '../index.js'
import Vector2 from '../modules/Vector2.js'
import Component from './Component.js'
import Modal from './Modals/Modal.js'

export default class ClickableGridHexagon extends Component {
	modal: Modal
	pos: Vector2
	size: number
	constructor(pos: Vector2, parentElement: HTMLElement, modal: Modal, size: number) {
		super('div', parentElement)
		this.pos = pos
		this.modal = modal
		this.size = size
		this.element.classList.add('clickableHexagon')
		this.addStyles({
			height: getPx(size),
			width: getPx(size),
			backgroundImage: "url('../../assets/Hexagon.png')",
			backgroundSize: getPx(size) + ' ' + getPx(size),
			transition: 'all .2s ease',
			position: 'relative',
			display: 'inline-block',
			cursor: 'pointer',
		})
	}
	onClick(e: MouseEvent) {
		this.modal.open()
	}
	onHover(e: MouseEvent, mouseIn: boolean) {
		if (mouseIn) {
			this.addStyles({
				transform: 'scale(1.2)',
				zIndex: '5',
			})
		} else {
			this.addStyles({
				transform: 'scale(1)',
				zIndex: '1',
			})
		}
	}
}
