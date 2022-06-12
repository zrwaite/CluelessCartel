import { game, getPx } from '../index.js'
import Component from './Component.js'
import TextSection from './TextSection.js'

export default class StatusBar extends Component {
	constructor(parentElement: HTMLElement) {
		super('section', parentElement)
		this.addStyles({
			height: getPx(20),
            width: getPx(660),
            borderRadius: getPx(4),
			border: getPx(3) + ' solid black',
			position: 'absolute',
			backgroundColor: 'var(--color1)',
            top: getPx(5),
            left: getPx(67)
		})

        let usernameText = new TextSection(this.element, 15, " " + game.user?.Username, {
            left: getPx(5),
            position: 'absolute',
            top: getPx(4)
        })
        let baseText = new TextSection(this.element, 15, "" + game.base?.Location.Name, {
            width: '100%',
            textAlign: 'center',
            position: 'absolute',
            top: getPx(4)
        })
        let cashText = new TextSection(this.element, 15, "$ " + game.user?.Cash, {
            right: getPx(5),
            position: 'absolute',
            top: getPx(4)
        })
	}
}
