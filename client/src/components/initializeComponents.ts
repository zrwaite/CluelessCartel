import { GameState } from '../game.js'
import { getPx, windowSize } from '../index.js'
import { navigate } from '../modules/navigate.js'
import AnimatedButton from './Buttons/AnimatedButton.js'
import AnimatorButton from './Buttons/AnimatorButton.js'
import Button from './Buttons/Button.js'
import OpenModalButton from './Buttons/OpenModalButton.js'
import ClickableGridHexagonRow from './ClickableGridHexagonRow.js'
import HexModal from './Modals/HexModal.js'
import SettingsModal from './Modals/SettingsModal.js'

const initializeSections = (gameElement: HTMLElement):{uiElement:HTMLElement, canvasElement:HTMLElement, hexElement: HTMLElement} => {
	let uiElement = document.createElement('div')
	uiElement.id = 'uiSection'
	gameElement.appendChild(uiElement)
	
	let canvasElement = document.createElement('div')
	canvasElement.id = 'canvasSection'
	uiElement.appendChild(canvasElement)
	
	let hexElement = document.createElement('div')
	hexElement.id = 'hexSection'
	canvasElement.appendChild(hexElement)
	
	;[gameElement, canvasElement, hexElement, uiElement].forEach((element) => {
		element.style.height = getPx(windowSize.y)
		element.style.width = getPx(windowSize.x)
		if (element !== gameElement) element.style.position = 'absolute'
	})
	return {uiElement, canvasElement, hexElement}
}

export const initializeStartComponents = (gameElement: HTMLElement, changeState: (newState: GameState) => void) => {
	let {uiElement} = initializeSections(gameElement)
	let startButton = new Button(uiElement, () => changeState('playing'), {
		top: `calc(50% - ${getPx(80)})`,
		left: `calc(50% - ${getPx(100)})`,
		height: getPx(60),
		width: getPx(200),
	})
	startButton.element.innerText = 'Start'

	let exitButton = new Button(uiElement, () =>navigate('/'), {
		top: `calc(50% + ${getPx(30)})`,
		left: `calc(50% - ${getPx(100)})`,
		height: getPx(60),
		width: getPx(200),
	})
	exitButton.element.innerText = 'Exit'
}

export const initializePlayComponents = (gameElement: HTMLElement, changeState: (newState: GameState) => void) => {
	let {uiElement, hexElement} = initializeSections(gameElement)
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

	let homeButton = new Button(uiElement, () => changeState('start'), {
		left: getPx(5),
		top: getPx(5),
	})
	homeButton.initializeIcon('home.svg')

	let storageButton = new Button(uiElement, () => alert('haha:)'), {
		right: getPx(5),
		bottom: getPx(5),
	})
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

export const clearComponents = (gameElement: HTMLElement) => {
	gameElement.innerHTML = ''
}
