import Button from './components/Buttons/Button.js'
import ClickableGridHexagon from './components/ClickableGridHexagon.js'
import ClickableGridHexagonRow from './components/ClickableGridHexagonRow.js'
import Modal from './components/Modal.js'
import { getPx } from './index.js'

export default class Game {
	editMenuOpen = false
	constructor() {}
	start(hexElement: HTMLElement, uiElement: HTMLElement) {}
	update() {
		// console.log("Hello");
	}
	draw(staticCtx: CanvasRenderingContext2D, dynamicCtx: CanvasRenderingContext2D) {}
}
