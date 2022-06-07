import { clearComponents, initializePlayComponents, initializeStartComponents } from './components/initializeComponents.js'
import { gameData } from './types/data.js'

export type GameState = 'start' | 'playing'

export default class Game {
	#state: GameState = 'start'
	element: HTMLElement
	data = gameData
	constructor(gameElement: HTMLElement) {
		this.element = gameElement
	}
	start() {
		clearComponents(this.element)
		if (this.#state === 'start') {
			initializeStartComponents(this.element, this)
		} else if (this.#state === 'playing') {
			initializePlayComponents(this.element, this)
		}
	}
	update() {
		if (this.#state === 'start') return
	}
	draw(staticCtx: CanvasRenderingContext2D, dynamicCtx: CanvasRenderingContext2D) {
		if (this.#state === 'start') return
	}
	changeState(newState: GameState) {
		this.#state = newState
		this.start()
	}
}
