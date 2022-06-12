import { getPx } from '../index.js'
import { StyleObject } from '../types/styles.js'
import Component from './Component.js'
export default class TextSection extends Component {
	constructor(parentElement: HTMLElement, textSize: number, text: string, styles: StyleObject = {}) {
		super('p', parentElement)
		this.addStyles({
			fontSize: getPx(textSize),
			lineHeight: getPx(textSize),
			margin: '0',
			...styles
		})
		this.setText(text)
	}
	setText(text: string) {
		this.element.innerText = text
	}
}
