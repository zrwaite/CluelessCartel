export default class Modal {
	element:HTMLElement;
	constructor() {
		const parentElement = document.getElementsByTagName("body")[0];
		this.element = document.createElement("div");
		this.element.style.display = "none";
		this.element.style.position = "absolute";
		this.element.style.height = "500px";
		this.element.style.width = "500px";
		this.element.style.top = "50%";
		this.element.style.left = "50%";
		this.element.style.transform = "translate(-50%, -50%)";
		this.element.style.backgroundColor = "blue";
		parentElement.appendChild(this.element);
		this.element.style.zIndex = "10000";
	}
	open() {
		this.element.style.display = "block";
	}
	close() {
		this.element.style.display = "none";
	}
}