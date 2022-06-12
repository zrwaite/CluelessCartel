import { getPx } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import AnimatedButton from './AnimatedButton.js'
import Button from './Button.js'

export default class AnimatorButton extends Button {
	animated: boolean = false
	animationFunctions: ((a: boolean) => void)[] = []
	constructor(parentElement: HTMLElement, animationButtons: AnimatedButton[], icon: string, styles: StyleObject = {}) {
		super(parentElement,() => {
			// this.addStyles({
				
			// })
			this.animationFunctions.forEach((fn) => fn(!this.animated))
			this.animated = !this.animated
		}, styles)
		this.initializeIcon(icon)
		animationButtons.forEach((btn) => this.animationFunctions.push(btn.animate.bind(btn)))
	}
}
