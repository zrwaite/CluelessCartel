import { getPx } from '../index.js'
import Component from './Component.js'
export default class TextSection extends Component {
	element: HTMLParagraphElement
	constructor(parentElement: HTMLElement, textSize: number, text: string) {
		super('div', parentElement)
		this.element = document.createElement('p')
		this.addStyles({
			fontSize: getPx(textSize),
			lineHeight: getPx(textSize),
			margin: '0',
		})
		parentElement.appendChild(this.element)
		this.setText(text)
	}
	setText(text: string) {
		this.element.innerText = text
	}
}
