import Game from './game.js'

export const windowSize = {
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

let staticCanvas = document.createElement('canvas')
canvasElement.appendChild(staticCanvas)
let staticCtx = staticCanvas.getContext('2d')

let dynamicCanvas = document.createElement('canvas')
canvasElement.appendChild(dynamicCanvas)
let dynamicCtx = dynamicCanvas.getContext('2d')

const game = new Game(gameElement)

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

export {game}