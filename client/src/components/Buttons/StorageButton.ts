import { game, getPx } from '../../index.js'
import { StyleObject } from '../../types/styles.js'
import Button from './Button.js'
import TextSection from '../TextSection.js'
import Icon from '../Icon.js'
import { gameAPI } from '../../modules/gameAPI.js'

export default class StorageButton extends Button {
	buttonContent: HTMLElement
	expandContent: HTMLElement
	expanded = false
	
	constructor(parentElement: HTMLElement, styles: StyleObject = {}) {
		super(parentElement, () => {}, {
			transition: 'all 0.7s',
            right: getPx(5),
		    bottom: getPx(5),
            height: getPx(40),
			width: getPx(50),
			borderRadius: getPx(6),
			border: getPx(3) + ' solid black',
			position: 'absolute',
			backgroundColor: 'var(--color1)',
			cursor: 'pointer',
			padding: '0',
			...styles,
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
        
		this.initializeButtonContent()	
		this.initializeExpandedContent()
		this.element.onmouseenter = (e) => this.onHover(e, true)
		this.element.onmouseleave = (e) => this.onHover(e, false)
	}

	initializeButtonContent() {
		let storageIcon = new Icon(this.buttonContent, 'storage.svg', {
			height: getPx(40),
			width: getPx(40),
			transition: 'all 0.7s'
		})
	}
	async buyBase() {
		const response = await gameAPI('/base', 'POST', {
			username: game.user.Username,
			function: 'new',
	
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
		game.base?.Resources.forEach(resource => {
			let resourceText = new TextSection(this.expandContent, 15, resource.ResourceName + ": " + resource.Amount)
		})

		

		let MethText = new TextSection(this.expandContent, 15, "Meth: " + game.base?.Drugs.Meth)
		let opiodsText = new TextSection(this.expandContent, 15, "Opiods: " + game.base?.Drugs.Opioids)
		let weedText = new TextSection(this.expandContent, 15, "Weed: " + game.base?.Drugs.Weed)

		let gunsText = new TextSection(this.expandContent, 15, "Guns: " + game.base?.Weapons.Guns)
		let explosivesText = new TextSection(this.expandContent, 15, "Explosives: " + game.base?.Weapons.Explosives)
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
			this.addStyles({
				height: getPx(170),
				width: getPx(120),
				border: `solid ${getPx(3)} grey`,
			})
			this.buttonContent.style.filter = 'opacity(0)'
		} else {
			this.expandContent.style.display = "none"
			this.expandContent.style.opacity = "0"
			setTimeout(() => {
				if (!this.expanded) {
					this.buttonContent.style.display = "block"
					this.buttonContent.style.filter = 'opacity(1)'
				}
			}, 700)
			this.addStyles({
				height: getPx(40),
				width: getPx(50),
				border: getPx(3) + ' solid black',
			})
		}
		this.expanded = !this.expanded
	}
	onHover(e:MouseEvent, mouseIn: boolean) {
		this.animate(mouseIn)
	}

}