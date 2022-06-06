import { clearComponents, initializePlayComponents, initializeStartComponents } from './components/initializeComponents.js'

export type GameState = 'start' | 'playing'

export default class Game {
	#state: GameState = 'start'
	element: HTMLElement
	constructor(gameElement: HTMLElement) {
		this.element = gameElement
	}
	start() {
		if (this.#state === 'start') {
			initializeStartComponents(this.element, this.changeState.bind(this))
		} else if (this.#state === 'playing') {
			initializePlayComponents(this.element, this.changeState.bind(this))
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
		clearComponents(this.element)
		this.start()
	}
}
