import { getPx } from '../index.js'
import { Hexagon } from '../types/hexagon.js'
import Component from './Component.js'
import HexagonModal from './Modals/HexagonModal.js'
import Modal from './Modals/Modal.js'

export default class ClickableGridHexagon extends Component {
	height = 115.5
	width = 100
	texture?: HTMLImageElement
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
			this.texture = document.createElement('img') as HTMLImageElement
			let imageLink
			if (hexagon.LandMaterial === 'Dirt') imageLink = 'DirtHexagon'
			if (hexagon.LandMaterial === 'Pavement') imageLink = 'PavementHexagon'
			if (hexagon.LandMaterial === 'Grass') imageLink = 'GrassHexagon'
			if (hexagon.LandMaterial === 'Sand') imageLink = 'SandHexagon'
			this.texture.src = `url('../../../../assets/${imageLink}.png`
			this.texture.style.height = '100%'
			this.texture.style.width = '101%'
			if (!hexagon.Owned) {
				this.texture.style.filter = 'grayscale(70%) brightness(70%)'
			}
			this.element.appendChild(this.texture)
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
