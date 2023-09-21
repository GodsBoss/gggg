//go:build js && wasm

package dominteraction

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"

	"syscall/js"
)

func FromMouseEvent(typ mouse.EventType, domEvent js.Value) mouse.Event {
	target := domEvent.Get("target")
	return mouse.Event{
		Type:   typ,
		Alt:    domEvent.Get("altKey").Bool(),
		Ctrl:   domEvent.Get("ctrlKey").Bool(),
		Shift:  domEvent.Get("shiftKey").Bool(),
		X:      domEvent.Get("clientX").Int() - target.Get("offsetLeft").Int(),
		Y:      domEvent.Get("clientY").Int() - target.Get("offsetTop").Int(),
		Button: mouse.Button(domEvent.Get("button").Int()),
	}
}
