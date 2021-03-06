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
	id: string
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
			else if (hexagon.LandMaterial.Name === 'Pavement') imageLink = 'PavementHexagon'
			else if (hexagon.LandMaterial.Name === 'Grass') imageLink = 'GrassHexagon'
			else if (hexagon.LandMaterial.Name === 'Sand') imageLink = 'SandHexagon'
			else if (hexagon.LandMaterial.Name === 'Water') imageLink = 'WaterHexagon'
			else throw Error('Invalid land material: <' + hexagon.LandMaterial.Name + '>')

			this.texture = new Icon(this.element, `${imageLink}.png`, {
				position: 'absolute',
				height: '100%',
				width: '101%',
				transform: `rotate(${hexagon.Rotation * 60}deg)`,
				left: '0',
				top: '0',
			})
			if (!hexagon.Owned) {
				this.texture.addStyles({
					filter: 'grayscale(70%) brightness(70%)',
				})
			}
			
			let { filename, found } = getStructureImage(hexagon.Structure.Name)
			if (found) {
				let structureImage = new Icon(this.element, filename, {
					position: 'absolute',
					width: '70%',
					height: '70%',
					left: '15%',
					top: '15%',
				})
				if (!hexagon.Owned) {
					structureImage.addStyles({
						filter: 'grayscale(70%) brightness(70%)',
					})
				}
			} else {
				if (hexagon.Structure.Name !== 'Empty') {
					let structure = new TextSection(this.element, 20, hexagon.Structure.Name)
					structure.addStyles({
						position: 'absolute',
						top: '20%',
						left: '20%',
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
