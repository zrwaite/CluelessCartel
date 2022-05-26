import Component from "./Component.js";
export default class Icon extends Component {
	element: HTMLImageElement
    constructor (parentElement: HTMLElement, assetFilename:string) {
		super("div", parentElement);
		this.element = new Image();
		this.element.src = "../../assets/"+assetFilename
		this.addStyles({
			height: "100%",
			width: "100%",
		})
		parentElement.appendChild(this.element)
	}
}