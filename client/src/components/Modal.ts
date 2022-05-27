import { getPx } from "../index.js";
import Button from "./Button.js";
import ClickableGridHexagon from "./ClickableGridHexagon.js";
import Component from "./Component.js";
import Icon from "./Icon.js";

export default class Modal extends Component {
	closeButton:Button
	text:HTMLElement
	constructor() {
		super("div",  document.getElementsByTagName("body")[0]);
		
		this.text = document.createElement("p")
		this.element.appendChild(this.text);
		this.addStyles({
			display: "none",
			position: "absolute",
			height: getPx(500),
			width: getPx(650),
			top: "50%",
			left: "50%",
			transform: "translate(-50%, -50%)",
			borderRadius: getPx(15),
			border: getPx(3) + " solid black",
			backgroundColor: "#994C00",
			zIndex: "10000",
		})
		this.closeButton = new Button(this.element, this.close.bind(this))
		this.closeButton.addStyles({
			height: getPx(30), 
			width: getPx(38), 
			right: getPx(-4),
			top: getPx(-4),
			border: getPx(3) + " solid black",
			borderRadius: getPx(6),
			position: "absolute",
			backgroundColor: "red",
		});
		this.closeButton.initializeIcon("close.svg")
	}
	open(hex:ClickableGridHexagon) {
		this.addStyles({display: "block"});
		this.text.innerText = `${hex.pos.x},${hex.pos.y}`
		

	}
	close() {
		this.addStyles({display: "none"});
		this.text.innerText = ""
	}
}