import Button from "./components/Button.js";
import ClickableGridHexagonRow from "./components/ClickableGridHexagonRow.js";
import Modal from "./components/Modal.js";
import {getPx} from "./index.js"

export default class Game {
	editMenuOpen = false
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
			right: getPx(5),
			top: getPx(5),
		});
		settingsButton.initializeIcon("settings.svg")

		let homeButton = new Button(uiElement,  ()=>{alert("ahhhhh fuck oh fuck oh shit")});
		homeButton.addStyles({
			left: getPx(5),
			top: getPx(5),
		});
		homeButton.initializeIcon("home.svg")

		let storageButton = new Button(uiElement,  ()=>{alert("haha:)")});
		storageButton.addStyles({
			right: getPx(5),
			bottom: getPx(5),
		});
		storageButton.initializeIcon("storage.svg")

		let buildButton = new Button(uiElement,  ()=>{alert("is that erin over there")});
		buildButton.addStyles({
			left: getPx(5),
			bottom: getPx(5),
			transition: "bottom 0.7s"
		});
		buildButton.initializeIcon("build.svg")
		buildButton.animate = (animateIn: boolean) => {
			if (animateIn) buildButton.addStyles({bottom: getPx(140)})
			else  buildButton.addStyles({bottom: getPx(5)})
		}

		let moveButton = new Button(uiElement,  ()=>{alert("is that erin over there")});
		moveButton.addStyles({
			left: getPx(5),
			bottom: getPx(5),
			transition: "bottom 0.7s"
		});
		moveButton.initializeIcon("move.svg")
		moveButton.animate = (animateIn: boolean) => {
			if (animateIn) moveButton.addStyles({bottom: getPx(95)})
			else  moveButton.addStyles({bottom: getPx(5)})
		}

		let deleteButton = new Button(uiElement,  ()=>{alert("is that erin over there")});
		deleteButton.addStyles({
			left: getPx(5),
			bottom: getPx(5),
			transition: "bottom 0.7s"
		});
		deleteButton.initializeIcon("delete.svg")
		deleteButton.animate = (animateIn: boolean) => {
			if (animateIn) deleteButton.addStyles({bottom: getPx(50)})
			else  deleteButton.addStyles({bottom: getPx(5)})
		}

		let editButton = new Button(uiElement,  ()=>{
			deleteButton.animate(!this.editMenuOpen)
			moveButton.animate(!this.editMenuOpen)
			buildButton.animate(!this.editMenuOpen)
			this.editMenuOpen = !this.editMenuOpen;
		})
			
		editButton.addStyles({
			left: getPx(5),
			bottom: getPx(5),
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