import ClickableGridHexagonRow from "./components/ClickableGridHexagonRow.js";
import Modal from "./components/Modal.js";

export default class Game {
	constructor() {

	}
	start(rootElement: HTMLElement) {
		let modal = new Modal();
		for (let i1=0; i1<10; i1++) {
			let newrow = new ClickableGridHexagonRow(i1, rootElement);
			for (let i2=0; i2<10; i2++) {
				newrow.addHexagon(i2);
				newrow.hexagons[i2].onClick = () => modal.open();
			}
		let modal = new Modal();
		}
		// setTimeout(() => modal.open(), 2000);
	}
	update() {
		// console.log("Hello"); 
	}
	draw(staticCtx: CanvasRenderingContext2D, dynamicCtx: CanvasRenderingContext2D) {

	}
}