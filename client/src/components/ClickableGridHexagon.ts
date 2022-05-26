import Vector2 from "../modules/Vector2.js";
import Component from "./Component.js";
import Modal from "./Modal.js";

export default class ClickableGridHexagon extends Component {
	modal: Modal;
	pos: Vector2;
	constructor(pos: Vector2, parentElement: HTMLElement, modal: Modal) {
		super("div", parentElement);
		this.pos = pos;
		this.modal = modal;
		this.element.classList.add("clickableHexagon");
		this.addStyles({
			height: "100px",
			width: "100px",
			backgroundImage: "url('../../assets/Hexagon.png')",
			backgroundSize: "100px 100px",
			transition: "all .2s ease",
			position: "relative",
			display: "inline-block",
			cursor: "pointer"
		})
	}
	onClick (e:MouseEvent) {
		this.modal.open(this);
	}
	onHover(e:MouseEvent, mouseIn:boolean){
		if (mouseIn){
			this.addStyles({
				transform: "scale(1.2)",
				zIndex: "5",
			})
		} else{
			this.addStyles({
				transform: "scale(1)",
				zIndex: "1",
			})
		}
	}
}