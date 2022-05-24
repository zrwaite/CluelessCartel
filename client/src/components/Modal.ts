import Component from "./Component.js";

export default class Modal extends Component {
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
			backgroundColor: "blue",
			zIndex: "10000",
		})
	}
	open() {
		this.addStyles({display: "block"});
	}
	close() {
		this.addStyles({display: "none"});
	}
}