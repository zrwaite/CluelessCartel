import { game, getPx, unit } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Button from '../Buttons/Button.js'
import ClickableGridHexagonRow from '../ClickableGridHexagonRow.js'
import Modal from './Modal.js'

export default class NewBaseModal extends Modal {
	constructor(backGround: string, styles: StyleObject = {}) {
		super('settings', styles)
		// const hexElement = document.getElementById('hexSection')
		// if (!hexElement) throw Error('hexSection not found')
        this.initializeBackground(backGround)

		
	}
}

