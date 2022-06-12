import { getPx } from '../index.js'
import { getStructureImage } from '../modules/getImageName.js'
import { Hexagon } from '../types/hexagon.js'
import Component from './Component.js'
import Icon from './Icon.js'
import HexagonModal from './Modals/HexagonModal.js'
import Modal from './Modals/Modal.js'
import TextSection from './TextSection.js'

export default class ClickableGridHexagon extends Component {
	height = 115.5
	width = 100
	texture?: Icon
	hexagon: Hexagon
	modal?: Modal
	id:string
	constructor(parentElement: HTMLElement, hexagon: Hexagon) {
		super('div', parentElement)
		this.hexagon = hexagon
		this.id = `button-${hexagon.X}-${hexagon.Y}`
		this.element.id = this.id
		this.element.classList.add('clickableHexagon')
		this.addStyles({
			height: getPx(this.height),
			width: getPx(this.width),
			transition: 'transform .2s ease',
			position: 'relative',
			display: 'inline-block',
			cursor: 'pointer',
		})
		if (hexagon.Buyable || hexagon.Owned) {
			this.modal = new HexagonModal(hexagon)
			this.element.classList.add('modalButton')
			let imageLink
			if (hexagon.LandMaterial.Name === 'Dirt') imageLink = 'DirtHexagon'
			if (hexagon.LandMaterial.Name === 'Pavement') imageLink = 'PavementHexagon'
			if (hexagon.LandMaterial.Name === 'Grass') imageLink = 'GrassHexagon'
			if (hexagon.LandMaterial.Name === 'Sand') imageLink = 'SandHexagon'
			if (hexagon.LandMaterial.Name === 'Water') imageLink = 'WaterHexagon'
			let degrees = Math.round(Math.random()*6) * 60
			this.texture = new Icon(this.element, `${imageLink}.png`, {
				position: 'absolute',
				height: '100%',
				width: '101%',
				transform: `rotate(${degrees}deg)`,
				left: '0',
				top: '0'
			})
			if (!hexagon.Owned) {
				this.texture.addStyles({
					filter: 'grayscale(70%) brightness(70%)'
				})
			} 
			let {filename, found} = getStructureImage(hexagon.Structure.Name)
			if (found) {
				let degrees = Math.round(Math.random()*6) * 60
				let structureImage = new Icon(this.element, filename, {
					position: 'absolute',
					width: getPx(50),
					height: getPx(50),
					left: getPx(25),
					top: getPx(32),
					transform: `rotate(${degrees}deg)`,
				})
			} else {
				if (hexagon.Structure.Name !== "Empty") {
					let structure = new TextSection(this.element, 20, hexagon.Structure.Name)
					structure.addStyles({
						position: 'absolute',
						top: getPx(25),
						left: getPx(25)
					})
				}
			}
		}
	}
	onClick(e: MouseEvent) {
		this.modal?.open()
	}
	onHover(e: MouseEvent, mouseIn: boolean) {
		if (mouseIn) {
			this.addStyles({
				transform: 'scale(1.2)',
				zIndex: '5',
			})
		} else {
			this.addStyles({
				transform: 'scale(1)',
				zIndex: '1',
			})
		}
	}
}
