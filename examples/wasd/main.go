package main

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/maininit"
)

func main() {
	maininit.Run(&game{})
}

type game struct{}

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

func (g *game) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (g *game) Renderer() maininit.Renderer {
	return g
}

func (g *game) SetOutput(ctx2d *dom.Context2D) {}

func (g *game) Render() {}

func (g *game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int) {
	return 640, 400
}
