
import Vector2 from "./modules/Vector2.js";

class button {
    element:HTMLElement

    constructor (parentElement: HTMLElement, onClick:Function) {
		this.element = document.createElement("button");

		this.element.style.height = "10px";
		this.element.style.width = "10px";
		this.element.onclick = (e) => onClick(e);
		this.element.style.backgroundImage = "url('assets/Hexagon.png')";
		this.element.style.backgroundSize = "100px 100px";

		parentElement.appendChild(this.element);
		this.element.onmouseenter = (() => this.onHover(true));
		this.element.onmouseleave = (() => this.onHover(false));
		this.element.style.transition = "all .2s ease";
	}
    onClick () {
		
	}
    onHover(mouseIn:boolean){
		if (mouseIn){
			this.element.style.transform = "scale(1.2)"; //Scale doesn't work in chrome
			this.element.style.zIndex = "1000";
		}else{
			this.element.style.transform = "scale(1)"; //Scale doesn't work in chrome
			this.element.style.zIndex = "1";
		}
	}
}