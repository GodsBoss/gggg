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
	tps := game.Config().TicksPerSecond()
	if tps <= 0 || tps > 1000 {
		return errors.NewString("0 < tps <= 1000 is violated")
	}
	w, err := dom.GlobalWindow()
	if err != nil {
		return err
	}
	doc, err := w.Document()
	if err != nil {
		return err
	}
	canvas, err := doc.CreateCanvasElement()
	if err != nil {
		return err
	}
	gameWidth, gameHeight := game.Config().GraphicsSize()
	canvas.SetSize(gameWidth, gameHeight)
	gameElement, err := doc.GetElementByID("game")
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
		msPerTick := 1000 / game.Config().TicksPerSecond()
		ticker := time.NewTicker(time.Millisecond * time.Duration(msPerTick))
		for {
			<-ticker.C
			game.Loop().Tick(msPerTick)
		}
	}()

	// Setup render loop.
	var reqAnimationFrameCallback func()
	reqAnimationFrameCallback = func() {
		w.RequestAnimationFrame(reqAnimationFrameCallback)
		game.Renderer().Render()
	}
	w.RequestAnimationFrame(reqAnimationFrameCallback)

	return nil
}
