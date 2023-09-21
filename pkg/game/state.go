package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/event/tick"
)

// state represents a game state. It does not hold any game data by itself, instead the data is passed to the event handlers.
// When a state is created, all handlers will be set to be non-nil, so there is no need to check when calling them.
type state[Data any] struct {
	InitHandler              func(data Data) NextState
	ReceiveKeyEventHandler   func(data Data, event keyboard.Event) NextState
	ReceiveMouseEventHandler func(data Data, event mouse.Event) NextState
	ReceiveTickEventHandler  func(data Data, event tick.Event) NextState
}

// Init calls the init handler. This is called when the state is entered.
func (st state[Data]) Init(data Data) NextState {
	return st.InitHandler(data)
}

// ReceiveKeyEvent calls the handler for key events.
func (st state[Data]) ReceiveKeyEvent(data Data, event keyboard.Event) NextState {
	return st.ReceiveKeyEventHandler(data, event)
}

// ReceiveMouseEvent calls the handler for mouse events.
func (st state[Data]) ReceiveMouseEvent(data Data, event mouse.Event) NextState {
	return st.ReceiveMouseEventHandler(data, event)
}

// ReceiveTickEvent calls the handler for tick events, i.e. the game clock.
func (st state[Data]) ReceiveTickEvent(data Data, event tick.Event) NextState {
	return st.ReceiveTickEventHandler(data, event)
}

// NextState is the next state an instance should be switched to after receiving an event.
// The first return value represents the stte to switch to, the second one whether the state
// should be switched at all.
type NextState func() (StateID, bool)

// SameState returns a NextState function which does not switch the state.
func SameState() NextState {
	return func() (StateID, bool) {
		return nil, false
	}
}

// SwitchState returns a NextState function which switches the state.
func SwitchState(id StateID) NextState {
	return func() (StateID, bool) {
		return id, true
	}
}
