import Vector2 from "../modules/Vector2.js";
import ClickableGridHexagon from "./ClickableGridHexagon.js";
import Component from "./Component.js";

export default class ClickableGridHexagonRow extends Component {
	index: number;
	hexagons: ClickableGridHexagon[] = [];
	constructor(index: number, parentElement:HTMLElement) {
		super("div", parentElement);
		this.index = index;
		this.element.classList.add("clickableHexagonRow");
		if (index%2!==0) {
			this.addStyles({
				top: `-${29*index}px`,
				left: `50px`
			})
		} else if (index!==0) {
			this.addStyles({top: `-${29*index}px`})
		}
	}
	addHexagon(xIndex:number) {
		this.hexagons.push(new ClickableGridHexagon(new Vector2(xIndex,this.index), this.element));
	}
}