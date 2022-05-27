import Button from "./components/Button.js";
import ClickableGridHexagonRow from "./components/ClickableGridHexagonRow.js";
import Modal from "./components/Modal.js";
import {getPx} from "./index.js"

export default class Game {
	constructor() {
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
			height: getPx(40), 
			width: getPx(50), 
			right: getPx(5),
			top: getPx(5),
			border: getPx(3) + " solid black",
			borderRadius: getPx(6),
			position: "absolute",
			backgroundColor: "grey",
		});
		settingsButton.initializeIcon("settings.svg")

		let homeButton = new Button(uiElement,  ()=>{alert("ahhhhh fuck oh fuck oh shit")});
		homeButton.addStyles({
			height: getPx(40), 
			width: getPx(50), 
			left: getPx(5),
			top: getPx(5),
			border: getPx(3) + " solid black",
			borderRadius: getPx(6),
			position: "absolute",
			backgroundColor: "grey",
		});
		homeButton.initializeIcon("home.svg")

		let storageButton = new Button(uiElement,  ()=>{alert("haha:)")});
		storageButton.addStyles({
			height: getPx(40), 
			width: getPx(50), 
			right: getPx(5),
			bottom: getPx(5),
			border: getPx(3) + " solid black",
			borderRadius: getPx(6),
			position: "absolute",
			backgroundColor: "grey",
		});
		storageButton.initializeIcon("storage.svg")

		let editButton = new Button(uiElement,  ()=>{alert("haha:)")});
		editButton.addStyles({
			height: getPx(40), 
			width: getPx(50), 
			left: getPx(5),
			bottom: getPx(5),
			border: getPx(3) + " solid black",
			borderRadius: getPx(6),
			position: "absolute",
			backgroundColor: "grey",
		});
		editButton.initializeIcon("edit.svg")

		let buildButton = new Button(uiElement,  ()=>{alert("is that erin over there")});
		buildButton.addStyles({
			height: getPx(40), 
			width: getPx(50), 
			left: getPx(5),
			bottom: getPx(50),
			border: getPx(3) + " solid black",
			borderRadius: getPx(6),
			position: "absolute",
			backgroundColor: "grey",
		});
		buildButton.initializeIcon("build.svg")

		// setTimeout(() => modal.open(), 2000);
	}
	update() {
		// console.log("Hello"); 
	}
	draw(staticCtx: CanvasRenderingContext2D, dynamicCtx: CanvasRenderingContext2D) {

	}
}