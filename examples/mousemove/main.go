package main

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/maininit"
)

func main() {
	win, _ := dom.GlobalWindow()
	doc, _ := win.Document()
	sprite, _ := doc.CreateImageElement("../assets/small_square.png")
	maininit.Run(
		&game{
			sprite: sprite,
		},
	)
	<-make(chan struct{}, 0)
}

type game struct {
	sprite *dom.Image
	output *dom.Context2D

	x int
	y int
}

func (g *game) Config() maininit.Config {
	return maininit.SimpleConfig{
		TPS: 40,
	}
}

func (g *game) Logic() maininit.Logic {
	return g
}

func (g *game) Tick(ms int) {}

func (g *game) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (g *game) ReceiveMouseEvent(event interaction.MouseEvent) {
	g.x = event.X
	g.y = event.Y
}

func (g *game) Renderer() maininit.Renderer {
	return g
}

func (g *game) SetOutput(ctx2d *dom.Context2D) {
	g.output = ctx2d
}

func (g *game) Render() {
	g.output.ClearRect(0, 0, 640, 400)
	g.output.DrawImage(g.sprite, 0, 0, 20, 20, g.x-10, g.y-10, 20, 20)
}

func (g *game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, xScale, yScale float64) {
	return 640, 400, 1.0, 1.0
}
