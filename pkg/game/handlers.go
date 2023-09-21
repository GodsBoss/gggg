package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/event/tick"
)

// InitHandler of a state is called whenever the state is entered. Returns the next state.
type InitHandler[Data any] func(data Data) NextState

func initHandlerNop[Data any](data Data) NextState {
	return SameState()
}

// KeyEventHandler of a state is called whenever a key event is received. Returns the next state.
type KeyEventHandler[Data any] func(data Data, event keyboard.Event) NextState

func keyboardHandlerNop[Data any](data Data, event keyboard.Event) NextState {
	return SameState()
}

// MouseEventHandler of a state is called whenever a mouse event is received. Returns the next state.
type MouseEventHandler[Data any] func(data Data, event mouse.Event) NextState

func mouseHandlerNop[Data any](data Data, event mouse.Event) NextState {
	return SameState()
}

// TickEventHandler of a state is called whenever a tick occurs. Returns the next state.
type TickEventHandler[Data any] func(data Data, event tick.Event) NextState

func tickHandlerNop[Data any](data Data, event tick.Event) NextState {
	return SameState()
}
