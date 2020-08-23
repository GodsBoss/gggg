package main

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/dominit"
	"github.com/GodsBoss/gggg/pkg/interaction"
)

func main() {
	win, _ := dom.GlobalWindow()
	doc, _ := win.Document()
	sprite, _ := doc.CreateImageElement("../assets/small_square.png")
	dominit.Run(
		&game{
			sprite: sprite,
		},
	)
	<-make(chan struct{}, 0)
}

type game struct {
	sprite *dom.Image
	output *dom.Context2D

	left  int
	right int
	down  int
	up    int
}

func (g *game) TicksPerSecond() int {
	return 40
}

func (g *game) Logic() dominit.Logic {
	return g
}

func (g *game) Tick(ms int) {}

func (g *game) ReceiveKeyEvent(event interaction.KeyEvent) {
	if event.Type == interaction.KeyDown {
		switch event.Key {
		case "w":
			g.up = 1
		case "a":
			g.left = 1
		case "s":
			g.down = 1
		case "d":
			g.right = 1
		}
	}
	if event.Type == interaction.KeyUp {
		switch event.Key {
		case "w":
			g.up = 0
		case "a":
			g.left = 0
		case "s":
			g.down = 0
		case "d":
			g.right = 0
		}
	}
}

func (g *game) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (g *game) Renderer() dominit.Renderer {
	return g
}

func (g *game) SetOutput(ctx2d *dom.Context2D) {
	g.output = ctx2d
}

func (g *game) Render() {
	g.output.ClearRect(0, 0, 640, 400)
	x := 320 + (g.right-g.left)*160 - 10
	y := 200 + (g.down-g.up)*100 - 10
	g.output.DrawImage(g.sprite, 0, 0, 20, 20, x, y, 20, 20)
}

func (g *game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, xScale, yScale float64) {
	return 640, 400, 1.0, 1.0
}
