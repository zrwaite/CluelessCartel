import { getPx } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Modal from '../Modal.js'
import Button from './Button.js'

export default class OpenModalButton extends Button {
	constructor(parentElement: HTMLElement, modal: Modal, icon: string, styles: StyleObject = {}) {
		super(parentElement, modal.open.bind(modal), styles)
		this.initializeIcon(icon)
	}
}
