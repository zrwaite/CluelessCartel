import Vector2 from "./modules/Vector2.js";

class ClickableHexagonRow {
	index: number;
	rowDiv: HTMLElement;
	hexagons: ClickableHexagon[] = [];
	constructor(index: number, parentElement: HTMLElement) {
		this.index = index;
		this.rowDiv = document.createElement("div");
		this.rowDiv.classList.add("clickableHexagonRow");
		console.log(parentElement.offsetHeight);
		if (index%2!==0) {
			this.rowDiv.style.top = `-${29*index}px`;
			this.rowDiv.style.left = `50px`;
		} else if (index!==0) {
			this.rowDiv.style.top = `-${29*index}px`;
		}
		parentElement.appendChild(this.rowDiv);
	}
	addHexagon(xIndex:number) {
		this.hexagons.push(new ClickableHexagon(new Vector2(xIndex,this.index), this.rowDiv));
	}
}

class ClickableHexagon {
	pos: Vector2;
	hexButton: HTMLElement;
	constructor (pos: Vector2, parentElement: HTMLElement) {
		this.pos = pos;
		this.hexButton = document.createElement("div");
		this.hexButton.classList.add("clickableHexagon");
		this.hexButton.style.height = "100px";
		this.hexButton.style.width = "100px";
		this.hexButton.onclick = (() => this.onClick());
		this.hexButton.style.backgroundImage = "url('assets/Hexagon.png')";
		this.hexButton.style.backgroundSize = "100px 100px";
		parentElement.appendChild(this.hexButton);
		this.hexButton.onmouseenter = (() => this.onHover(true));
		this.hexButton.onmouseleave = (() => this.onHover(false));
	}
	onClick () {
		alert(`${this.pos.x},${this.pos.y}`);
	}
	onHover(mouseIn:boolean){
		if (mouseIn){
			this.hexButton.style.scale = "1.2"; //Scale doesn't work in chrome
			this.hexButton.style.zIndex = "1000";
		}else{
			this.hexButton.style.scale = "1"
			this.hexButton.style.zIndex = "1";
		}
	}
}

export {ClickableHexagon, ClickableHexagonRow}