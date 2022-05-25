import Game from "./game.js";

const WIDTH=800;
const HEIGHT=500; 

let rootElement = document.getElementById("gameSection");
if (!rootElement) throw Error("Element #gameSection does not exist");

let canvasHolder = document.getElementById("canvasHolder");
if (!canvasHolder) throw Error("Element #canvasHolder does not exist");
canvasHolder.style.height = "100%";
canvasHolder.style.width = "100%";

[canvasHolder, rootElement].forEach(element => {
	element.style.height = `${HEIGHT}px`;
	element.style.width = `${WIDTH}px`;
});

let staticCanvas = document.createElement("canvas");
canvasHolder.appendChild(staticCanvas);
let staticCtx = staticCanvas.getContext("2d");

let dynamicCanvas = document.createElement("canvas");
canvasHolder.appendChild(dynamicCanvas);
let dynamicCtx = dynamicCanvas.getContext("2d");

let game = new Game();
game.start(rootElement); 

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
