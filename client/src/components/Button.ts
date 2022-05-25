import Component from "./Component.js";
import Icon from "./Icon.js";
export default class Button extends Component {
	icon: Icon;
    constructor (parentElement: HTMLElement, onClick:Function) {
		super("button", parentElement);
		this.addStyles({
			height: "10px",
			width: "10px",
		})
		this.onClick = (e:MouseEvent) => onClick(e)
		this.icon = new Icon(this.element, ()=>{});
	}
    onHover(e: MouseEvent, mouseIn:boolean){
		if (mouseIn){
			this.element.style.transform = "scale(1.1)"; //Scale doesn't work in chrome
			this.element.style.zIndex = "1000";
		}else{
			this.element.style.transform = "scale(1)"; //Scale doesn't work in chrome
			this.element.style.zIndex = "1";
		}
	}
}