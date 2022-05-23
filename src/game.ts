import {ClickableHexagon, ClickableHexagonRow} from "./clickableHexagon.js";
import Vector2 from "./modules/Vector2.js";

export default class Game {
	constructor() {

	}
	start(rootElement: HTMLElement) {
		for (let i1=0; i1<10; i1++) {
			let newrow = new ClickableHexagonRow(i1, rootElement);
			for (let i2=0; i2<10; i2++) {
				newrow.addHexagon(i2)
			}
		}
	}
	update() {
		// console.log("Hello"); 
	}
	draw(staticCtx: CanvasRenderingContext2D, dynamicCtx: CanvasRenderingContext2D) {

	}
}