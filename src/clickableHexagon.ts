import Vector2 from "./modules/Vector2.js";

export default class ClickableHexagon {
	pos: Vector2;
	hexButton: HTMLElement;
	constructor (pos: Vector2, parentElement: HTMLElement) {
		this.pos = pos;
		this.hexButton = document.createElement("div");
		this.hexButton.classList.add("clickableHexagon");
		this.hexButton.style.height = "100px";
		this.hexButton.style.width = "100px";
		if (pos.y%2!==0) {
			this.hexButton.style.top = `-${29*pos.y}px`;
			this.hexButton.style.left = `50px`;
		} else if (pos.y!==0) {
			this.hexButton.style.top = `-${29*pos.y}px`;
		}
		this.hexButton.onclick = (() => this.onClick());
		this.hexButton.style.backgroundImage = "url('assets/Hexagon.png')";
		this.hexButton.style.backgroundSize = "100px 100px";
		parentElement.appendChild(this.hexButton);

		this.hexButton.onmouseenter = (() => this.onHover(true))
		this.hexButton.onmouseleave = (() => this.onHover(false))
	}
	onClick () {
		alert(`${this.pos.x},${this.pos.y}`);
	}
	onHover(mouseIn:boolean){
		if (mouseIn){
			console.log(mouseIn);
			this.hexButton.style.scale = "1.2"
			this.hexButton.style.zIndex = "1000";
		}else{
			this.hexButton.style.scale = "1"
			this.hexButton.style.zIndex = "1";
		}
	}
}