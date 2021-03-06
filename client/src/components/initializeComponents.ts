import { game, getPx, windowSize } from '../index.js'
import { navigate } from '../modules/navigate.js'
import AnimatedButton from './Buttons/AnimatedButton.js'
import AnimatorButton from './Buttons/AnimatorButton.js'
import Button from './Buttons/Button.js'
import OpenModalButton from './Buttons/OpenModalButton.js'
import ClickableGridHexagonRow from './ClickableGridHexagonRow.js'
import Modal from './Modals/Modal.js'
import SettingsModal from './Modals/SettingsModal.js'
import TextSection from './TextSection.js'
import NewBaseModal from './Modals/NewBaseModal.js'
import StorageButton from './Buttons/StorageButton.js'
import StatusBar from './StatusBar.js'

const initializeSections = (gameElement: HTMLElement): { uiElement: HTMLElement; canvasElement: HTMLElement; hexElement: HTMLElement } => {
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
	return { uiElement, canvasElement, hexElement }
}

export const initializeStartComponents = (gameElement: HTMLElement) => {
	let { uiElement } = initializeSections(gameElement)
	let startButton = new Button(uiElement, () => game.changeState('bases'), {
		top: `calc(50% - ${getPx(80)})`,
		left: `calc(50% - ${getPx(100)})`,
		height: getPx(60),
		width: getPx(200),
	})
	new TextSection(startButton.element, 30, 'Start')

	let exitButton = new Button(uiElement, () => navigate('/'), {
		top: `calc(50% + ${getPx(30)})`,
		left: `calc(50% - ${getPx(100)})`,
		height: getPx(60),
		width: getPx(200),
	})
	new TextSection(exitButton.element, 30, 'Exit')
}

export const initializeBasesComponents = (gameElement: HTMLElement) => {
	let { uiElement } = initializeSections(gameElement)
	let y = 0
	game.user.Bases.forEach((base, i) => {
		y = Math.round(i / 4)
		let baseButton = new Button(
			uiElement,
			() => {
				game.base = base
				game.changeState('playing')
			},
			{
				top: getPx(100 + 60 * y),
				left: getPx(5 + 90 * i),
				height: getPx(50),
				width: getPx(80),
			}
		)
		new TextSection(baseButton.element, 10, base.Location.Name)
	})
	let newBaseModal = new NewBaseModal('MapOfUSA.svg')
	let newBaseButton = new OpenModalButton(uiElement, newBaseModal, 'add.svg', {
		top: getPx(100 + 60 * (y + 1)),
		left: getPx(5),
		height: getPx(50),
		width: getPx(80),
	})
	new TextSection(newBaseButton.element, 10, 'New Base')
}

export const initializePlayComponents = (gameElement: HTMLElement) => {
	document.getElementsByTagName('body')[0].onclick = (e: MouseEvent) => {
		if (e.target) {
			let element = e.target as HTMLElement
			if (!element.closest(`.modal`) && !element.closest('.modalButton')) {
				game.modals.forEach((modal) => modal.close())
			}
		}
	}

	let { uiElement, hexElement } = initializeSections(gameElement)
	let hexagonRows: ClickableGridHexagonRow[] = []

	let statusBar = new StatusBar(uiElement)

	let base = game.user.Bases.find(thisBase => thisBase.Location.Name === game.base?.Location.Name)
	if (!base) {
		alert("Base not found")
		throw Error("Base not found")
	}

	base.HexagonRows.forEach((hexagonRow) => {
		let newrow = new ClickableGridHexagonRow(hexElement, hexagonRow)
		newrow.initializeHexagons()
		hexagonRows.push(newrow)
	})

	let settingsModal = new SettingsModal(hexagonRows)
	let settingsButton = new OpenModalButton(uiElement, settingsModal, 'settings.svg', {
		right: getPx(5),
		top: getPx(5),
	})

	hexElement.scrollTop = game.scroll.y
	hexElement.scrollLeft = game.scroll.x

	let homeButton = new Button(uiElement, () => game.changeState('start'), {
		left: getPx(5),
		top: getPx(5),
	})
	homeButton.initializeIcon('home.svg')

	let storageButton = new StorageButton(uiElement, )

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
