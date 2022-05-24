import Vector2 from "../modules/Vector2.js";
import Component from "./Component.js";

export default class ClickableGridHexagon extends Component {
	pos: Vector2;
	constructor(pos: Vector2, parentElement: HTMLElement) {
		super("div", parentElement);
		this.pos = pos;
		this.element.classList.add("clickableHexagon");
		this.addStyles({
			height: "100px",
			width: "100px",
			backgroundImage: "url('assets/Hexagon.png')",
			backgroundSize: "100px 100px",
			transition: "all .2s ease"
		})
	}
	onClick (e:MouseEvent) {
		alert(`${this.pos.x},${this.pos.y}`);
	}
	onHover(e:MouseEvent, mouseIn:boolean){
		if (mouseIn){
			this.addStyles({
				transform: "scale(1.2)",
				zIndex: "1000",
			})
		} else{
			this.addStyles({
				transform: "scale(1)",
				zIndex: "1",
			})
		}
	}
}