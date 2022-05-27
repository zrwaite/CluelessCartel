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
			width: "50px", 
			right: "5px",
			top: "5px",
			border: "3px solid black",
			borderRadius: "6px",
			position: "absolute",
			backgroundColor: "grey",
		});
		settingsButton.initializeIcon("settings.svg")

		let homeButton = new Button(uiElement,  ()=>{alert("ahhhhh fuck oh fuck oh shit")});
		homeButton.addStyles({
			height: "40px", 
			width: "50px", 
			left: "5px",
			top: "5px",
			border: "3px solid black",
			borderRadius: "6px",
			position: "absolute",
			backgroundColor: "grey",
		});
		homeButton.initializeIcon("home.svg")

		let storageButton = new Button(uiElement,  ()=>{alert("haha:)")});
		storageButton.addStyles({
			height: "40px", 
			width: "50px", 
			right: "5px",
			bottom: "5px",
			border: "3px solid black",
			borderRadius: "6px",
			position: "absolute",
			backgroundColor: "grey",
		});
		storageButton.initializeIcon("storage.svg")

		let editButton = new Button(uiElement,  ()=>{alert("haha:)")});
		editButton.addStyles({
			height: "40px", 
			width: "50px", 
			left: "5px",
			bottom: "5px",
			border: "3px solid black",
			borderRadius: "6px",
			position: "absolute",
			backgroundColor: "grey",
		});
		editButton.initializeIcon("edit.svg")

		// setTimeout(() => modal.open(), 2000);
	}
	update() {
		// console.log("Hello"); 
	}
	draw(staticCtx: CanvasRenderingContext2D, dynamicCtx: CanvasRenderingContext2D) {

	}
}