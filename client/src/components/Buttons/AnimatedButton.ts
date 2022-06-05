import { getPx } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Modal from '../Modal.js'
import Button from './Button.js'

export default class AnimatedButton extends Button {
	start: number
	end: number
	constructor(parentElement: HTMLElement, onClick: Function, icon: string, start: number, end: number, styles: StyleObject = {}) {
		super(parentElement, onClick, {
			transition: 'bottom 0.7s',
			...styles,
		})
		this.initializeIcon(icon)
		this.start = start
		this.end = end
	}
	animate(animateIn: boolean) {
		if (animateIn) this.addStyles({ bottom: getPx(this.end) })
		else this.addStyles({ bottom: getPx(this.start) })
	}
}
