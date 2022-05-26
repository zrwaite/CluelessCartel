import Game from "./game.js";

const WIDTH=800;
const HEIGHT=500; 

let hexElement = document.getElementById("hexSection");
if (!hexElement) throw Error("Element #hexSection does not exist");

let uiElement = document.getElementById("uiSection");
if (!uiElement) throw Error("Element #uiElement does not exist");

let canvasElement = document.getElementById("canvasSection");
if (!canvasElement) throw Error("Element #canvasSection does not exist");
canvasElement.style.height = "100%";
canvasElement.style.width = "100%";

[canvasElement, hexElement, uiElement].forEach(element => {
	element.style.height = `${HEIGHT}px`;
	element.style.width = `${WIDTH}px`;
	element.style.position = "absolute";
});

let staticCanvas = document.createElement("canvas");
canvasElement.appendChild(staticCanvas);
let staticCtx = staticCanvas.getContext("2d");

let dynamicCanvas = document.createElement("canvas");
canvasElement.appendChild(dynamicCanvas);
let dynamicCtx = dynamicCanvas.getContext("2d");

let game = new Game();
game.start(hexElement, uiElement); 

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
