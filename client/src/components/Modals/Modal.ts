import { game, getPx } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Button from '../Buttons/Button.js'
import Component from '../Component.js'

export default class Modal extends Component {
	closeButton: Button
	id:string
	constructor(id:string, styles: StyleObject = {}) {
		const parentElement = document.getElementById('gameSection')
		if (!parentElement) throw Error('#gameSection not found')
		super('div', parentElement)
		this.id = id
		this.element.id = `modal-${this.id}`
		this.element.classList.add('modal')
		this.addStyles({
			display: 'none',
			position: 'absolute',
			height: getPx(350),
			width: getPx(410),
			top: '50%',
			left: '50%',
			transform: 'translate(-50%, -50%)',
			borderRadius: getPx(15),
			border: getPx(3) + ' solid silver',
			backgroundColor: 'var(--color5)',
			//backgroundcolor: (0,0,0), /* Fallback color */
			//backgroundcolor: rgba(0,0,0,0.9), /* Black w/ opacity */
			zIndex: '10000',
			...styles,
		})
		this.closeButton = new Button(this.element, this.close.bind(this), {
			height: getPx(30),
			width: getPx(38),
			right: getPx(-4),
			top: getPx(-4),
			border: getPx(3) + ' silver',
			borderRadius: getPx(6),
			position: 'absolute',
			backgroundColor: 'red',
		})
		this.closeButton.initializeIcon('close.svg')
		game.modals.push(this)
	}
	open() {
		game.modals.forEach((modal) => modal.close())
		this.addStyles({ display: 'block' })
	}
	close() {
		this.addStyles({ display: 'none' })
	}
}
