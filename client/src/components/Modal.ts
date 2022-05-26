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
			height: "500px",
			width: "650px",
			top: "50%",
			left: "50%",
			transform: "translate(-50%, -50%)",
			borderRadius: "15px",
			border: "3px solid black",
			backgroundColor: "#994C00",
			zIndex: "10000",
		})
		this.closeButton = new Button(this.element, this.close.bind(this))
		this.closeButton.addStyles({
			height: "30px", 
			width: "30px", 
			right: "-4px",
			top: "-4px",
			border: "3px solid black",
			borderRadius: "6px",
			position: "absolute",
			backgroundColor: "red",
		});
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