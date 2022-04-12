//go:build js && wasm

package canvas2drendering

import (
	"github.com/GodsBoss/gggg/pkg/dom"
)

func MonochromeBackground(color string) (RenderFunc, error) {
	fillStyle, err := dom.NewColorCanvasFillStyle(color)
	if err != nil {
		return nil, err
	}
	return func(ctx *dom.Context2D) {
		ctx.SetFillStyle(fillStyle)
		w, h := ctx.Size()
		ctx.FillRect(0, 0, w, h)
	}, nil
}
