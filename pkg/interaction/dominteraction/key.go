//go:build js && wasm

package dominteraction

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"

	"syscall/js"
)

func FromKeyEvent(typ keyboard.EventType, domEvent js.Value) keyboard.Event {
	return keyboard.Event{
		Type:  typ,
		Alt:   domEvent.Get("altKey").Bool(),
		Ctrl:  domEvent.Get("ctrlKey").Bool(),
		Shift: domEvent.Get("shiftKey").Bool(),
		Key:   domEvent.Get("key").String(),
	}
}
