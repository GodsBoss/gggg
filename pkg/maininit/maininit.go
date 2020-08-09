// build js,wasm

package maininit

import (
	"syscall/js"
	"time"

	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/errors"
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/interaction/dominteraction"
)

// Game represents the game as a whole - configuration, rendering and logic.
type Game interface {
	Config() Config
	Renderer() Renderer
	Logic() Logic
}

// Config are a few parameters needed to setup the game.
type Config interface {
	TicksPerSecond() int
}

// Renderer renders the game into a given output.
type Renderer interface {
	SetOutput(ctx2d *dom.Context2D)
	Render()

	// Scale takes the available size and returns the real size of the canvas the renderer
	// will perform on.
	Scale(availableWidth, availableHeight int) (realWidth, realHeight int)
}

// Logic is the game logic.
type Logic interface {
	Tick(ms int)
	ReceiveKeyEvent(event interaction.KeyEvent)
	ReceiveMouseEvent(event interaction.MouseEvent)
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
	resizeCanvas(game.Renderer(), window, canvas)
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
	go runGameLoop(1000/ticksPerSecond, game.Logic())

	// Setup render loop.
	runRendering(window, game.Renderer())

	passGameEvents(game.Logic(), window, canvas)

	return nil
}

func runGameLoop(msPerTick int, logic Logic) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(msPerTick))
	for {
		<-ticker.C
		logic.Tick(msPerTick)
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

func passGameEvents(logic Logic, window *dom.Window, canvas *dom.Canvas) {
	dom.AddEventListener(
		window,
		"keydown",
		func(event js.Value) {
			logic.ReceiveKeyEvent(dominteraction.FromKeyEvent(interaction.KeyDown, event))
		},
	)
	dom.AddEventListener(
		window,
		"keyup",
		func(event js.Value) {
			logic.ReceiveKeyEvent(dominteraction.FromKeyEvent(interaction.KeyUp, event))
		},
	)
	dom.AddEventListener(
		canvas,
		"mousedown",
		func(event js.Value) {
			logic.ReceiveMouseEvent(dominteraction.FromMouseEvent(interaction.MouseDown, event))
		},
	)
	dom.AddEventListener(
		canvas,
		"mouseup",
		func(event js.Value) {
			logic.ReceiveMouseEvent(dominteraction.FromMouseEvent(interaction.MouseUp, event))
		},
	)
	dom.AddEventListener(
		canvas,
		"mousemove",
		func(event js.Value) {
			logic.ReceiveMouseEvent(dominteraction.FromMouseEvent(interaction.MouseMove, event))
		},
	)
}

func resizeCanvas(renderer Renderer, window *dom.Window, canvas *dom.Canvas) {
	canvas.SetSize(renderer.Scale(window.InnerSize()))
}
