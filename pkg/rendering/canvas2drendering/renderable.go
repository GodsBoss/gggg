// +build js,wasm

package canvas2drendering

import (
	"github.com/GodsBoss/gggg/pkg/dom"
)

// Renderable is something that can be rendered to a DOM 2D canvas context.
type Renderable interface {
	// Render renders this onto the canvas.
	Render(ctx *dom.Context2D)
}

// Renderables is a list of renderables, which are rendered in order.
type Renderables []Renderable

func (rs Renderables) Render(ctx *dom.Context2D) {
	for i := range rs {
		rs[i].Render(ctx)
	}
}

// RenderFunc is a function implementation of Renderable.Render.
type RenderFunc func(*dom.Context2D)

func (f RenderFunc) Render(ctx *dom.Context2D) {
	f(ctx)
}
