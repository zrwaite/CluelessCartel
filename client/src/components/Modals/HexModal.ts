import { StyleObject } from '../../types/styles.js'
import Modal from './Modal.js'

export default class HexModal extends Modal {
	constructor(parentElement: HTMLElement, styles: StyleObject = {}) {
		super(parentElement, styles)
	}
}
