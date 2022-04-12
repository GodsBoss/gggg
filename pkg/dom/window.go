//go:build js && wasm

package dom

import (
	"github.com/GodsBoss/gggg/v2/pkg/errors"

	"syscall/js"
)

// Window is a wrapper for the global window property of browers. The zero value
// is not useful, create an instance with NewWindow() or GlobalWindow().
type Window struct {
	value js.Value
}

// NewWindow creates a Window instance.
func NewWindow(jsWindow js.Value) (*Window, error) {
	if jsWindow.IsNull() {
		return nil, errors.NewString("window object must not be null")
	}
	return &Window{
		value: jsWindow,
	}, nil
}

// GlobalWindow creates a Window instance by taking the global value.
func GlobalWindow() (*Window, error) {
	return NewWindow(js.Global())
}

func (w *Window) getValue() js.Value {
	v := w.value
	if !v.IsNull() {
		return v
	}
	return js.Global()
}

func (w *Window) getJSNode() js.Value {
	return w.value
}

func (w *Window) Document() (*Document, error) {
	jsDoc := w.getValue().Get("document")
	if jsDoc.IsNull() {
		return nil, errors.NewString("document object does not exist")
	}
	return &Document{
		value: jsDoc,
	}, nil
}

func (w *Window) RequestAnimationFrame(f func()) {
	w.getValue().Call(
		"requestAnimationFrame",
		js.FuncOf(
			func(_ js.Value, _ []js.Value) interface{} {
				f()
				return nil
			},
		),
	)
}

func (w *Window) InnerSize() (width int, height int) {
	return w.getValue().Get("innerWidth").Int(), w.getValue().Get("innerHeight").Int()
}
