import { game, getPx } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Button from './Button.js'
import TextSection from '../TextSection.js'
import { gameAPI } from '../../modules/gameAPI.js'

export default class ExpandButton extends Button {
	buttonContent: HTMLElement
	expandContent: HTMLElement
	expanded = false
	title: TextSection
	name: string
	
	
	constructor(parentElement: HTMLElement, name: string, styles: StyleObject = {}) {
		super(parentElement, () => {}, {
			transition: 'all 0.7s',
			borderRadius: getPx(7.5),
			backgroundColor: "rgba(255, 0, 0, 0.5)",
			height: getPx(11),
			width: getPx(11),
			border: "solid 1px red",
			...styles,
		})
		this.name = name

		this.title = new TextSection(this.element, 15, name)
		this.title.addStyles({
			paddingLeft: getPx(11),
			transition: 'all 0.7s',
		})

		this.buttonContent = document.createElement("div")
		this.expandContent = document.createElement("div")
		this.expandContent.style.display = "none"
		;[this.buttonContent, this.expandContent].forEach(element => {
			element.style.height = "100%"
			element.style.width = "100%"
			element.style.transition = 'all 0.7s',
			this.element.appendChild(element)
		})
        
		this.initializeButtonContent(name)	
		this.initializeExpandedContent()
		this.element.onmouseenter = (e) => this.onHover(e, true)
		this.element.onmouseleave = (e) => this.onHover(e, false)
	}

	initializeButtonContent(text:string) {
	}
	async buyBase() {
		const response = await gameAPI('/base', 'POST', {
			username: game.user.Username,
			function: 'new',
			location: this.name,
		})
		if (response) {
			if (response.Success) {
				game.user = response.Response
			} else {
				alert(JSON.stringify(response.Errors))
			}
		} else throw Error('Failed to get user!')
		game.trySaveScroll()
		game.start()
	}
	initializeExpandedContent() {
		let BuyButton = new Button(this.expandContent, this.buyBase.bind(this), )
		BuyButton.addStyles({
			top: getPx(52),
			left: getPx(52),
		})
        let BuyButtonText = new TextSection(BuyButton.element, 17, "Buy")
        
	}
	animate(animateIn: boolean) {
		if (animateIn) {
			setTimeout(() => {
				if (this.expanded) {
					this.buttonContent.style.display = "none"
					this.expandContent.style.display = "block"
					this.expandContent.style.opacity = "1"
				}
			}, 700)
			this.title.addStyles({paddingLeft: getPx(0)})
			this.addStyles({
				height: getPx(100),
				width: getPx(150),
				border: `solid ${getPx(3)} grey`,
			})
		} else {
			this.expandContent.style.display = "none"
			this.expandContent.style.opacity = "0"
			setTimeout(() => {
				if (!this.expanded) {
					this.buttonContent.style.display = "block"
				}
			}, 700)
			this.title.addStyles({paddingLeft: getPx(11)})
			this.addStyles({
				height: getPx(11),
				width: getPx(11),
				border: `solid ${getPx(1)} red`,
			})
		}
		this.expanded = !this.expanded
	}
	onHover(e:MouseEvent, mouseIn: boolean) {
		this.animate(mouseIn)
	}

}