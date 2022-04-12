//go:build js && wasm

package dominteraction

import (
	"github.com/GodsBoss/gggg/v2/pkg/interaction"

	"syscall/js"
)

func FromKeyEvent(typ interaction.KeyEventType, domEvent js.Value) interaction.KeyEvent {
	return interaction.KeyEvent{
		Type:  typ,
		Alt:   domEvent.Get("altKey").Bool(),
		Ctrl:  domEvent.Get("ctrlKey").Bool(),
		Shift: domEvent.Get("shiftKey").Bool(),
		Key:   domEvent.Get("key").String(),
	}
}
