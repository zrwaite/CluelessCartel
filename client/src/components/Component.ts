import { StyleObject } from '../types/styles'

export default class Component {
	element: HTMLElement
	constructor(elementType: string, parentElement: HTMLElement) {
		this.element = document.createElement(elementType)
		this.element.onclick = (e) => this.onClick(e)
		parentElement.appendChild(this.element)
		this.element.onmouseenter = (e) => this.onHover(e, true)
		this.element.onmouseleave = (e) => this.onHover(e, false)
	}
	addStyles(styleObj: StyleObject) {
		for (const [key, value] of Object.entries(styleObj) as unknown as [number, string][]) {
			this.element.style[key] = value
		}
	}
	onClick(e: MouseEvent) {}
	onHover(e: MouseEvent, hoverIn: boolean) {}
}
export type ComponentArgs = [string, HTMLElement]
