//go:build js && wasm

package dom

import (
	"errors"
	"syscall/js"
)

// Canvas wraps a JS canvas element. Canvas implements Node.
type Canvas struct {
	value js.Value
}

func (canvas *Canvas) getJSNode() js.Value {
	return canvas.value
}

func (canvas *Canvas) SetSize(width, height int) {
	canvas.SetWidth(width)
	canvas.SetHeight(height)
}

func (canvas *Canvas) SetWidth(width int) {
	canvas.value.Set("width", width)
}

func (canvas *Canvas) SetHeight(height int) {
	canvas.value.Set("height", height)
}

func (canvas *Canvas) Context2D() (*Context2D, error) {
	jsCtx := canvas.value.Call("getContext", "2d")
	if jsCtx.IsNull() {
		return nil, errors.New("2d context not supported")
	}
	return &Context2D{
		value: jsCtx,
	}, nil
}

type Context2D struct {
	value js.Value
}

func (ctx2D *Context2D) DisableImageSmoothing() {
	ctx2D.value.Set("imageSmoothingEnabled", false)
}

func (ctx2D *Context2D) DrawImage(image *Image, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight int) {
	ctx2D.value.Call("drawImage", image.value, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight)
}

func (ctx2D *Context2D) ClearRect(x, y, width, height int) {
	ctx2D.value.Call("clearRect", x, y, width, height)
}

func (ctx2D *Context2D) Size() (w int, h int) {
	canv := ctx2D.value.Get("canvas")
	return canv.Get("width").Int(), canv.Get("height").Int()
}

func (ctx2D *Context2D) SetFillStyle(fillStyle CanvasFillStyle) {
	ctx2D.value.Set("fillStyle", fillStyle.color())
}

func (ctx2D *Context2D) FillRect(x, y, w, h int) {
	ctx2D.value.Call("fillRect", x, y, w, h)
}

// TODO: Support other fill styles.

type CanvasFillStyle interface {
	color() string
}

// NewColorCanvasFillStyle creates a monochrome canvas fill style. If the color
// value is invalid, an error is returned (not implemented yet).
func NewColorCanvasFillStyle(color string) (CanvasFillStyle, error) {
	return colorCanvasFillStyle(color), nil
}

type colorCanvasFillStyle string

func (style colorCanvasFillStyle) color() string {
	return string(style)
}
