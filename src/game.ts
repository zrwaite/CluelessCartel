import ClickableHexagon from "./clickableHexagon.js";
import Vector2 from "./modules/Vector2.js";

export default class Game {
	constructor() {

	}
	start(rootElement: HTMLElement) {

		for (let i1=0; i1<10; i1++) {
			let newrow = document.createElement("div");
			rootElement.appendChild(newrow)
			for (let i2=0; i2<10; i2++) {
				// let y = i1%2===0? 0:1;
				let firstHexagon = new ClickableHexagon(new Vector2(i2,i1), newrow)
			}
		}
	}
	update() {
		// console.log("Hello"); 
	}
	draw(staticCtx: CanvasRenderingContext2D, dynamicCtx: CanvasRenderingContext2D) {

	}
}