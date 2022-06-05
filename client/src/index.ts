import { initializeComponents } from './components/initializeComponents.js'
import Game from './game.js'

const windowSize = {
	// Accessible through game.width and game.height.
	x: 800,
	y: 500,
	border: 10,
}
export const unit =
	(window.innerWidth - windowSize.border * 2) / (window.innerHeight - windowSize.border * 2) > windowSize.x / windowSize.y
		? (window.innerHeight - windowSize.border * 2) / windowSize.y
		: (window.innerWidth - windowSize.border * 2) / windowSize.x
export const getPx = (num: number): string => (num * unit).toString() + 'px'

let gameElement = document.getElementById('gameSection')
if (!gameElement) throw Error('Element #gameSection does not exist')

let hexElement = document.getElementById('hexSection')
if (!hexElement) throw Error('Element #hexSection does not exist')

let uiElement = document.getElementById('uiSection')
if (!uiElement) throw Error('Element #uiElement does not exist')

let canvasElement = document.getElementById('canvasSection')
if (!canvasElement) throw Error('Element #canvasSection does not exist')
;[gameElement, canvasElement, hexElement, uiElement].forEach((element) => {
	element.style.height = getPx(windowSize.y)
	element.style.width = getPx(windowSize.x)
	if (element !== gameElement) element.style.position = 'absolute'
})

let staticCanvas = document.createElement('canvas')
canvasElement.appendChild(staticCanvas)
let staticCtx = staticCanvas.getContext('2d')

let dynamicCanvas = document.createElement('canvas')
canvasElement.appendChild(dynamicCanvas)
let dynamicCtx = dynamicCanvas.getContext('2d')

let game = new Game()
game.start(hexElement, uiElement)
initializeComponents(uiElement, hexElement)

let lastTime = 0
function gameLoop(timestamp: number) {
	if (!dynamicCtx || !staticCtx) throw Error('Context failed')
	let deltaTime = timestamp - lastTime
	lastTime = timestamp
	game.update()
	game.draw(staticCtx, dynamicCtx)
	requestAnimationFrame(gameLoop)
}
requestAnimationFrame(gameLoop)
