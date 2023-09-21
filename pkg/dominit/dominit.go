//go:build js && wasm

package dominit

import (
	"errors"
	"syscall/js"
	"time"

	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/interaction/dominteraction"
)

// Game represents the game as a whole - configuration, rendering and logic.
type Game interface {
	Config
	Renderer
	Logic
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
	// will perform on. It also returns a scaling factor.
	Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64)
}

// Logic is the game logic.
type Logic interface {
	Tick(ms int)
	ReceiveKeyEvent(event keyboard.Event)
	ReceiveMouseEvent(event mouse.Event)
}

func Run(game Game) {
	_ = run(game)
}

func run(game Game) error {
	ticksPerSecond := game.TicksPerSecond()
	if ticksPerSecond <= 0 || ticksPerSecond > 1000 {
		return errors.New("0 < ticksPerSecond <= 1000 is violated")
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
	mouseXYScaler := &mouseEventCoordinateScaler{}
	rw, rh, xf, yf := game.Scale(window.InnerSize())
	canvas.SetSize(rw, rh)
	mouseXYScaler.setScaling(xf, yf)
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
	game.SetOutput(context2D)

	// Setup game loop.
	go runGameLoop(1000/ticksPerSecond, game)

	// Setup render loop.
	runRendering(window, game)

	passGameEvents(game, window, canvas, mouseXYScaler)

	dom.AddEventListener(
		window,
		"resize",
		func(_ js.Value) {
			rw, rh, xf, yf := game.Scale(window.InnerSize())
			canvas.SetSize(rw, rh)
			mouseXYScaler.setScaling(xf, yf)
		},
	)

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

func passGameEvents(logic Logic, window *dom.Window, canvas *dom.Canvas, scaler *mouseEventCoordinateScaler) {
	dom.AddEventListener(
		window,
		"keydown",
		func(event js.Value) {
			logic.ReceiveKeyEvent(dominteraction.FromKeyEvent(keyboard.Down, event))
		},
	)
	dom.AddEventListener(
		window,
		"keyup",
		func(event js.Value) {
			logic.ReceiveKeyEvent(dominteraction.FromKeyEvent(keyboard.Up, event))
		},
	)
	dom.AddEventListener(
		canvas,
		"mousedown",
		func(event js.Value) {
			logic.ReceiveMouseEvent(scaler.scale(dominteraction.FromMouseEvent(mouse.Down, event)))
		},
	)
	dom.AddEventListener(
		canvas,
		"mouseup",
		func(event js.Value) {
			logic.ReceiveMouseEvent(scaler.scale(dominteraction.FromMouseEvent(mouse.Up, event)))
		},
	)
	dom.AddEventListener(
		canvas,
		"mousemove",
		func(event js.Value) {
			logic.ReceiveMouseEvent(scaler.scale(dominteraction.FromMouseEvent(mouse.Move, event)))
		},
	)
}

type mouseEventCoordinateScaler struct {
	xFactor float64
	yFactor float64
}

func (scaler *mouseEventCoordinateScaler) scale(ev mouse.Event) mouse.Event {
	ev.X = int(float64(ev.X) / scaler.xFactor)
	ev.Y = int(float64(ev.Y) / scaler.yFactor)
	return ev
}

func (scaler *mouseEventCoordinateScaler) setScaling(x, y float64) {
	scaler.xFactor = x
	scaler.yFactor = y
}
