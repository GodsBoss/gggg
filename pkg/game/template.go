package game

import (
	"errors"
)

// Template is a template for game instances.
type Template[Data any] struct {
	states        map[StateID]*stateTemplate[Data]
	initalStateID StateID

	// CreateData creates the data for a game instance. Optional.
	CreateData func() Data
}

// NewInstance creates a new game instance from a game template.
// At least one state must have been added to be successful.
func (tmpl Template[Data]) NewInstance() (*Instance[Data], error) {
	if len(tmpl.states) == 0 {
		return nil, errors.New("template does not contain any states")
	}

	var data Data
	if tmpl.CreateData != nil {
		data = tmpl.CreateData()
	}

	instance := &Instance[Data]{
		states: make(map[StateID]state[Data]),
		data:   data,
	}

	for id := range tmpl.states {
		instance.states[id] = tmpl.states[id].toInstanceState()
	}

	instance.nextState(SwitchState(tmpl.initalStateID))

	return instance, nil
}

// AddState adds a state. If this the first state that's added, the state is also the start state.
// Instances created by NewInstance() are not affected by adding states.
func (tmpl *Template[Data]) AddState() StateTemplate[Data] {
	st := &stateTemplate[Data]{
		template: tmpl,
		id:       newStateID(),
	}

	if tmpl.initalStateID == nil {
		tmpl.initalStateID = st.ID()
	}

	if tmpl.states == nil {
		tmpl.states = make(map[StateID]*stateTemplate[Data])
	}

	tmpl.states[st.ID()] = st

	return st
}

// StateTemplate represents the template for a state. When creating an instance from a game template,
// these states are taken as they are at the point of creation. Subsequently changing states does not
// affect instances already created, only future instances.
type StateTemplate[Data any] interface {
	// SetAsStartState makes this state the start state of the corresponding game template.
	// Overrides a previously set start state. Returns itself.
	SetAsStartState() StateTemplate[Data]

	// ID returns this state's ID.
	ID() StateID

	// SetInitHandler sets the init handler for this state. Returns itself.
	SetInitHandler(handler InitHandler[Data]) StateTemplate[Data]

	// SetKeyboardHandler sets the keyboard handler for this state. Returns itself.
	SetKeyboardHandler(handler KeyEventHandler[Data]) StateTemplate[Data]

	// SetMouseHandler sets the mouse handler for this state. Returns itself.
	SetMouseHandler(handler MouseEventHandler[Data]) StateTemplate[Data]

	// SetTickHandler sets the tick handler for this state. Returns itself.
	SetTickHandler(handler TickEventHandler[Data]) StateTemplate[Data]
}

type stateTemplate[Data any] struct {
	template *Template[Data]
	id       StateID

	initHandler     InitHandler[Data]
	keyboardHandler KeyEventHandler[Data]
	mouseHandler    MouseEventHandler[Data]
	tickHandler     TickEventHandler[Data]
}

func (tmpl *stateTemplate[Data]) SetAsStartState() StateTemplate[Data] {
	tmpl.template.initalStateID = tmpl.ID()
	return tmpl
}

func (tmpl *stateTemplate[Data]) ID() StateID {
	return tmpl.id
}

func (tmpl *stateTemplate[Data]) SetInitHandler(handler InitHandler[Data]) StateTemplate[Data] {
	tmpl.initHandler = handler
	return tmpl
}

func (tmpl *stateTemplate[Data]) SetKeyboardHandler(handler KeyEventHandler[Data]) StateTemplate[Data] {
	tmpl.keyboardHandler = handler
	return tmpl
}

func (tmpl *stateTemplate[Data]) SetMouseHandler(handler MouseEventHandler[Data]) StateTemplate[Data] {
	tmpl.mouseHandler = handler
	return tmpl
}

func (tmpl *stateTemplate[Data]) SetTickHandler(handler TickEventHandler[Data]) StateTemplate[Data] {
	tmpl.tickHandler = handler
	return tmpl
}

func (tmpl stateTemplate[Data]) toInstanceState() state[Data] {
	st := state[Data]{
		InitHandler:              initHandlerNop[Data],
		ReceiveKeyEventHandler:   keyboardHandlerNop[Data],
		ReceiveMouseEventHandler: mouseHandlerNop[Data],
		ReceiveTickEventHandler:  tickHandlerNop[Data],
	}

	if tmpl.initHandler != nil {
		st.InitHandler = tmpl.initHandler
	}

	if tmpl.keyboardHandler != nil {
		st.ReceiveKeyEventHandler = tmpl.keyboardHandler
	}

	if tmpl.mouseHandler != nil {
		st.ReceiveMouseEventHandler = tmpl.mouseHandler
	}

	if tmpl.tickHandler != nil {
		st.ReceiveTickEventHandler = tmpl.tickHandler
	}

	return st
}
