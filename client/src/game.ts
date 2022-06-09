import { clearComponents, initializeBasesComponents, initializePlayComponents, initializeStartComponents } from './components/initializeComponents.js'
import Modal from './components/Modals/Modal.js'
import { gameAPI } from './modules/gameAPI.js'
import { Base } from './types/base.js'
import { gameData } from './types/data.js'

export type GameState = 'start' | 'playing' | 'bases'

export default class Game {
	zoom: number = 1
	scroll: {x:number, y:number} = {x:0,y:0}
	#state: GameState = 'start'
	element: HTMLElement
	data = gameData
	base?: Base
	modals: Modal[] = []
	constructor(gameElement: HTMLElement) {
		this.element = gameElement
		this.initializeData()
	}
	async initializeData() {
		const response:any = await gameAPI('/user?username=Insomnizac')
		if (response && response.Success) {
			this.data = response.Response
			this.start()
		} else {
			this.start()
			console.error("Not in true game mode")
			// throw Error("Failed to get user!")
		}
	}
	start() {
		clearComponents(this.element)
		switch(this.#state) {
			case 'start':
				initializeStartComponents(this.element)
				break
			case 'playing':
				if (this.base) {
					initializePlayComponents(this.element)
					break
				}
			case 'bases': 
				initializeBasesComponents(this.element)
				break;
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
	trySaveScroll() {
		let hexElement = document.getElementById("hexSection") 
		if (hexElement) {
			this.scroll.x = hexElement.scrollLeft
			this.scroll.y = hexElement.scrollTop
		}
	}
}
