import { getPx } from '../index.js'
import Component from './Component.js'
export default class TextSection extends Component {
	constructor(parentElement: HTMLElement, textSize: number, text: string) {
		super('p', parentElement)
		this.addStyles({
			fontSize: getPx(textSize),
			lineHeight: getPx(textSize),
			margin: '0',
		})
		this.setText(text)
	}
	setText(text: string) {
		this.element.innerText = text
	}
}
