import { getPx } from "../index.js";
import Component from "./Component.js";
import Icon from "./Icon.js";
export default class Button extends Component {
	icon: Icon | null = null;
    constructor (parentElement: HTMLElement, onClick:Function) {
		super("button", parentElement);
		this.addStyles({
			height: getPx(10),
			width: getPx(10),
			cursor: "pointer",
			transition: "all .2s ease",
			padding: "0"
		})
		this.onClick = (e:MouseEvent) => onClick(e)
	}
	initializeIcon ( assetFilename:string){
		this.icon = new Icon(this.element, assetFilename)
	}
    onHover(e: MouseEvent, mouseIn:boolean){
		if (mouseIn){
			this.element.style.transform = "scale(1.1)"; 
		}else{
			this.element.style.transform = "scale(1)"; 
		}
	}
}