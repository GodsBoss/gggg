// build js,wasm

package maininit

import (
	"time"

	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/errors"
)

// Game represents the game as a whole - configuration, rendering and logic.
type Game interface {
	Config() Config
	Renderer() Renderer
	Loop() Loop
}

// Config are a few parameters needed to setup the game.
type Config interface {
	GraphicsSize() (width int, height int)
	TicksPerSecond() int
}

// Renderer renders the game into a given output.
type Renderer interface {
	SetOutput(ctx2d *dom.Context2D)
	Render()
}

// Loop is the game logic.
type Loop interface {
	Tick(ms int)
}

func Run(game Game) {
	_ = run(game)
}

func run(game Game) error {
	ticksPerSecond := game.Config().TicksPerSecond()
	if ticksPerSecond <= 0 || ticksPerSecond > 1000 {
		return errors.NewString("0 < ticksPerSecond <= 1000 is violated")
	}
	window, err := dom.GlobalWindow()
	if err != nil {
		return err
	}
	document, err := window.Document()
	if err != nil {
		return err
	}
	canvas, err := document.CreateCanvasElement()
	if err != nil {
		return err
	}
	gameWidth, gameHeight := game.Config().GraphicsSize()
	canvas.SetSize(gameWidth, gameHeight)
	gameElement, err := document.GetElementByID("game")
	if err != nil {
		return err
	}
	err = gameElement.AppendChild(canvas)
	if err != nil {
		return err
	}
	context2D, err := canvas.Context2D()
	if err != nil {
		return err
	}
	game.Renderer().SetOutput(context2D)

	// Setup game loop.
	go runGameLoop(1000/ticksPerSecond, game.Loop())

	// Setup render loop.
	runRendering(window, game.Renderer())

	return nil
}

func runGameLoop(msPerTick int, loop Loop) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(msPerTick))
	for {
		<-ticker.C
		loop.Tick(msPerTick)
	}
}

func runRendering(window *dom.Window, renderer Renderer) {
	var reqAnimationFrameCallback func()
	reqAnimationFrameCallback = func() {
		window.RequestAnimationFrame(reqAnimationFrameCallback)
		renderer.Render()
	}
	window.RequestAnimationFrame(reqAnimationFrameCallback)
}
