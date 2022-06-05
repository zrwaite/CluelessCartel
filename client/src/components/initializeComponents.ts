import { getPx } from '../index.js'
import { navigate } from '../modules/navigate.js'
import AnimatedButton from './Buttons/AnimatedButton.js'
import AnimatorButton from './Buttons/AnimatorButton.js'
import Button from './Buttons/Button.js'
import OpenModalButton from './Buttons/OpenModalButton.js'
import ClickableGridHexagonRow from './ClickableGridHexagonRow.js'
import HexModal from './Modals/HexModal.js'
import SettingsModal from './Modals/SettingsModal.js'

export const initializeComponents = (uiElement: HTMLElement, hexElement: HTMLElement) => {
	let hexModal = new HexModal(uiElement)
	let hexagonRows = []
	for (let i1 = 0; i1 < 10; i1++) {
		let newrow = new ClickableGridHexagonRow(i1, hexElement)
		for (let i2 = 0; i2 < 10; i2++) {
			newrow.addHexagon(i2, hexModal)
		}
		hexagonRows.push(newrow)
	}

	let settingsModal = new SettingsModal(uiElement, hexagonRows)

	let settingsButton = new OpenModalButton(uiElement, settingsModal, 'settings.svg', {
		right: getPx(5),
		top: getPx(5),
	})

	let homeButton = new Button(
		uiElement,
		() => {
			navigate('/')
		},
		{
			left: getPx(5),
			top: getPx(5),
		}
	)
	homeButton.initializeIcon('home.svg')

	let storageButton = new Button(
		uiElement,
		() => {
			alert('haha:)')
		},
		{
			right: getPx(5),
			bottom: getPx(5),
		}
	)
	storageButton.initializeIcon('storage.svg')

	let buildButton = new AnimatedButton(
		uiElement,
		() => {
			alert('is that erin over there')
		},
		'build.svg',
		5,
		140,
		{
			left: getPx(5),
			bottom: getPx(5),
		}
	)

	let moveButton = new AnimatedButton(
		uiElement,
		() => {
			alert('is that erin over there')
		},
		'move.svg',
		5,
		95,
		{
			left: getPx(5),
			bottom: getPx(5),
		}
	)

	let deleteButton = new AnimatedButton(
		uiElement,
		() => {
			alert('is that erin over there')
		},
		'delete.svg',
		5,
		50,
		{
			left: getPx(5),
			bottom: getPx(5),
		}
	)

	let editButton = new AnimatorButton(uiElement, [deleteButton, moveButton, buildButton], 'edit.svg', {
		left: getPx(5),
		bottom: getPx(5),
	})
}
