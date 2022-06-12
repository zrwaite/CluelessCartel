import { game } from "../../index.js"
import { gameAPI } from "../gameAPI.js"

export const buyStructure = async (x:number, y:number|undefined, structureName: string) => {
	const response = await gameAPI('/structure', 'POST', {
		username: game.user.Username,
		function: 'buy',
		BaseLocation: game.base?.Location.Name,
		HexagonX: x,
		HexagonY: y,
		StructureName: structureName
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