import Game from "./game.js";

let canvasHolder = document.getElementById("canvasHolder");
if (!canvasHolder) throw Error("Element #canvasHolder does not exist");

let staticCanvas = document.createElement("canvas");
canvasHolder.appendChild(staticCanvas);
let staticCtx = staticCanvas.getContext("2d");

let dynamicCanvas = document.createElement("canvas");
canvasHolder.appendChild(dynamicCanvas);
let dynamicCtx = dynamicCanvas.getContext("2d");


let game = new Game();
game.start();

let lastTime = 0;
function gameLoop(timestamp: number) {
	if (!dynamicCtx || !staticCtx) throw Error("Context failed");
	let deltaTime = timestamp - lastTime;
	lastTime = timestamp;
	game.update();
	game.draw(staticCtx, dynamicCtx);
	requestAnimationFrame(gameLoop);
}
requestAnimationFrame(gameLoop);
