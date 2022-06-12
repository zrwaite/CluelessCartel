import Component from './Component.js'
import { StyleObject } from '../types/styles.js'
export default class Icon extends Component {
	element: HTMLImageElement
	constructor(parentElement: HTMLElement, assetFilename: string, styles: StyleObject = {}) {
		super('div', parentElement)
		this.element = new Image()
		this.element.src = '../../assets/' + assetFilename
		this.addStyles({
			height: '100%',
			width: '100%',
			...styles
		})
		parentElement.appendChild(this.element)
	}
}
