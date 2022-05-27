import { getPx } from "../index.js";
import Vector2 from "../modules/Vector2.js";
import ClickableGridHexagon from "./ClickableGridHexagon.js";
import Component from "./Component.js";
import Modal from "./Modal.js";

export default class ClickableGridHexagonRow extends Component {
	index: number;
	hexagons: ClickableGridHexagon[] = [];
	constructor(index: number, parentElement:HTMLElement) {
		super("div", parentElement);
		this.index = index;
		this.element.classList.add("clickableHexagonRow");
		this.addStyles({position: "relative"})
		if (index%2!==0) {
			this.addStyles({
				top: getPx(-29*index),
				left: getPx(50),
			})
		} else if (index!==0) {
			this.addStyles({top: getPx(-29*index)})
		}
	}
	addHexagon(xIndex:number, modal: Modal) {
		this.hexagons.push(new ClickableGridHexagon(new Vector2(xIndex,this.index), this.element, modal));
	}
}