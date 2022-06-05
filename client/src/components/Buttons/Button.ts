import { getPx } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Component from '../Component.js'
import Icon from '../Icon.js'
import Modal from '../Modal.js'

export default class Button extends Component {
	icon: Icon | null = null

	constructor(parentElement: HTMLElement, onClick: Function, styles: StyleObject = {}) {
		super('button', parentElement)
		this.addStyles({
			height: getPx(40),
			width: getPx(50),
			borderRadius: getPx(6),
			border: getPx(3) + ' solid black',
			position: 'absolute',
			backgroundColor: 'grey',
			cursor: 'pointer',
			transition: 'all .2s ease',
			padding: '0',
			...styles,
		})
		this.onClick = (e: MouseEvent) => onClick(e)
	}
	initializeIcon(assetFilename: string) {
		this.icon = new Icon(this.element, assetFilename)
	}
	onHover(e: MouseEvent, mouseIn: boolean) {
		if (mouseIn) {
			this.element.style.transform = 'scale(1.1)'
		} else {
			this.element.style.transform = 'scale(1)'
		}
	}
}
