import Button from "./components/Button.js";
import ClickableGridHexagonRow from "./components/ClickableGridHexagonRow.js";
import Modal from "./components/Modal.js";

export default class Game {
	constructor() {
		let user = {
			name: "Dave",
			password: "Hi"
		} 
		console.log("user: ", user);
	}
	start(hexElement: HTMLElement, uiElement:HTMLElement) {
		let modal = new Modal();
		for (let i1=0; i1<10; i1++) {
			let newrow = new ClickableGridHexagonRow(i1, hexElement);
			for (let i2=0; i2<10; i2++) {
				newrow.addHexagon(i2, modal);
			}
		}

		let settingsButton = new Button(uiElement,  ()=>{alert("hi")});
		settingsButton.addStyles({
			height: "40px", 
			width: "40px", 
			right: "0px",
			top: "0px",
			border: "3px solid black",
			borderRadius: "6px",
			position: "absolute",
			backgroundColor: "grey",
		});

		// setTimeout(() => modal.open(), 2000);
	}
	update() {
		// console.log("Hello"); 
	}
	draw(staticCtx: CanvasRenderingContext2D, dynamicCtx: CanvasRenderingContext2D) {

	}
}