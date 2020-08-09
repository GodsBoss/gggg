// build js,wasm

package maininit

import (
	"time"

	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/errors"
)

type Game interface {
	Config() Config
	Renderer() Renderer
	Loop() Loop
}

type Config interface {
	GraphicsSize() (width int, height int)
	TicksPerSecond() int
}

type Renderer interface {
	SetOutput(ctx2d *dom.Context2D)
	Render()
}

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
	go func() {
		msPerTick := 1000 / ticksPerSecond
		ticker := time.NewTicker(time.Millisecond * time.Duration(msPerTick))
		for {
			<-ticker.C
			game.Loop().Tick(msPerTick)
		}
	}()

	// Setup render loop.
	var reqAnimationFrameCallback func()
	reqAnimationFrameCallback = func() {
		window.RequestAnimationFrame(reqAnimationFrameCallback)
		game.Renderer().Render()
	}
	window.RequestAnimationFrame(reqAnimationFrameCallback)

	return nil
}
