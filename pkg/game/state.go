package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/event/tick"
)

// State represents a game state. It does not hold any game data by itself, instead the data is passed to the event handlers.
// All handlers are optional, e.g. they may be nil.
type State[Data any] struct {
	InitHandler              func(data Data)
	ReceiveKeyEventHandler   func(data Data, event keyboard.Event) NextState
	ReceiveMouseEventHandler func(data Data, event mouse.Event) NextState
	ReceiveTickEventHandler  func(data Data, event tick.Event) NextState
}

// Init calls the init handler. This is called when the state is entered.
func (state State[Data]) Init(data Data) {
	if state.InitHandler != nil {
		state.InitHandler(data)
	}
}

// ReceiveKeyEvent calls the handler for key events.
func (state State[Data]) ReceiveKeyEvent(data Data, event keyboard.Event) NextState {
	if state.ReceiveKeyEventHandler != nil {
		return state.ReceiveKeyEventHandler(data, event)
	}
	return SameState()
}

// ReceiveMouseEvent calls the handler for mouse events.
func (state State[Data]) ReceiveMouseEvent(data Data, event mouse.Event) NextState {
	if state.ReceiveMouseEventHandler != nil {
		return state.ReceiveMouseEventHandler(data, event)
	}
	return SameState()
}

// ReceiveTickEvent calls the handler for tick events, i.e. the game clock.
func (state State[Data]) ReceiveTickEvent(data Data, event tick.Event) NextState {
	if state.ReceiveTickEventHandler != nil {
		return state.ReceiveTickEventHandler(data, event)
	}
	return SameState()
}

type StateID int

type NextState func() (StateID, bool)

// SameState returns a NextState function which does not switch the state.
func SameState() NextState {
	return func() (StateID, bool) {
		return 0, false
	}
}

// SwitchState returns a NextState function which switches the state.
func SwitchState(id StateID) NextState {
	return func() (StateID, bool) {
		return id, true
	}
}
