import { game } from "../../index.js"
import { gameAPI } from "../gameAPI.js"

export const buyHexagon = async (x:number, y:number|undefined) => {
	const response = await gameAPI('/hexagon', 'POST', {
		username: game.user.Username,
		function: 'buy',
		BaseLocation: game.base?.Location.Name,
		HexagonX: x,
		HexagonY: y,
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