//go:build js && wasm

package dominteraction

import (
	"github.com/GodsBoss/gggg/v2/pkg/interaction"

	"syscall/js"
)

func FromMouseEvent(typ interaction.MouseEventType, domEvent js.Value) interaction.MouseEvent {
	target := domEvent.Get("target")
	return interaction.MouseEvent{
		Type:   typ,
		Alt:    domEvent.Get("altKey").Bool(),
		Ctrl:   domEvent.Get("ctrlKey").Bool(),
		Shift:  domEvent.Get("shiftKey").Bool(),
		X:      domEvent.Get("clientX").Int() - target.Get("offsetLeft").Int(),
		Y:      domEvent.Get("clientY").Int() - target.Get("offsetTop").Int(),
		Button: domEvent.Get("button").Int(),
	}
}
