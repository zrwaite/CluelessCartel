import Button from "./Button.js";
import Component from "./Component.js";
import Icon from "./Icon.js";

export default class Modal extends Component {
	closeButton:Button
	constructor() {
		super("div",  document.getElementsByTagName("body")[0]);
		this.addStyles({
			display: "none",
			position: "absolute",
			height: "500px",
			width: "500px",
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
	open() {
		this.addStyles({display: "block"});
	}
	close() {
		this.addStyles({display: "none"});
	}
}