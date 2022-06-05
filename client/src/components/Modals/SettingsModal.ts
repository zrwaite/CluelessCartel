import { getPx, unit } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Button from '../Buttons/Button.js'
import Modal from './Modal.js'

export default class SettingsModal extends Modal {
	zoom: number = 1
	constructor(parentElement: HTMLElement, styles: StyleObject = {}) {
		super(parentElement, styles)
		// const hexElement = document.getElementById('hexSection')
		// if (!hexElement) throw Error('hexSection not found')

		const zoomIn = () => {
			const scale = 1.1 / unit
			const hexRows = document.querySelectorAll('#hexSection>*') as unknown as HTMLElement[]
			const hexs = document.querySelectorAll('#hexSection>*>*') as unknown as HTMLElement[]

			hexRows.forEach((hr) => {
				hr.style.height = getPx(hr.offsetHeight * scale)
				if (hr.style.left != '') hr.style.left = getPx(parseFloat(hr.style.left.split('px')[0]) * scale)
			})
			hexs.forEach((hr) => {
				let newHeight = getPx(hr.offsetHeight * scale)
				hr.style.backgroundSize = newHeight + ' ' + newHeight
				hr.style.height = newHeight
				hr.style.width = newHeight
			})
			// hexRows.forEach((hr) => (hr.style.height = hr.style.height))
			// hexs.forEach((hr) => (hr.style.transform = `scale(${this.zoom}`))
		}
		const zoomInButton = new Button(this.element, zoomIn.bind(this), {
			top: getPx(5),
			left: getPx(5),
		})
		zoomInButton.initializeIcon('add.svg')

		const zoomOut = () => {
			const scale = 1 / 1.1 / unit
			const hexRows = document.querySelectorAll('#hexSection>*') as unknown as HTMLElement[]
			const hexs = document.querySelectorAll('#hexSection>*>*') as unknown as HTMLElement[]
			hexRows.forEach((hr) => {
				hr.style.height = getPx(hr.offsetHeight * scale)
				if (hr.style.left != '') hr.style.left = getPx(parseFloat(hr.style.left.split('px')[0]) * scale)
			})
			hexs.forEach((hr) => {
				let newHeight = getPx(hr.offsetHeight * scale)
				hr.style.backgroundSize = newHeight + ' ' + newHeight
				hr.style.height = getPx(hr.offsetHeight * scale)
				hr.style.width = getPx(hr.offsetHeight * scale)
			})
			// hexs.forEach((hr) => (hr.style.transform = `scale(${this.zoom}`))
		}
		const zoomOutButton = new Button(this.element, zoomOut.bind(this), {
			top: getPx(45),
			left: getPx(5),
		})
		zoomOutButton.initializeIcon('remove.svg')
	}
}
