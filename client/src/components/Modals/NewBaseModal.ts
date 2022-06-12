import { game, getPx, unit } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Button from '../Buttons/Button.js'
import ExpandButton from '../Buttons/BuyBaseButton.js'
import ClickableGridHexagonRow from '../ClickableGridHexagonRow.js'
import Modal from './Modal.js'

export default class NewBaseModal extends Modal {
	constructor(backGround: string, styles: StyleObject = {}) {
		super('settings', styles)
		
        this.initializeBackground(backGround)
		let FlordiaButton = new ExpandButton(this.element, 'Florida')
		FlordiaButton.addStyles({
			top: getPx(244),
			left: getPx(355)
		})

		let NewYorkButton = new ExpandButton(this.element, 'New York')
		NewYorkButton.addStyles({
			top: getPx(110),
			left: getPx(372)
		})

		let MinnesotaButton = new ExpandButton(this.element, 'Minnesota')
		MinnesotaButton.addStyles({
			top: getPx(97),
			left: getPx(215)
		})

		let NevadaButton = new ExpandButton(this.element, 'Nevada')
		NevadaButton.addStyles({
			top: getPx(155),
			left: getPx(105)
		})


	}
	
}

