package game_test

import (
	"testing"

	"github.com/GodsBoss/gggg/v2/pkg/event/tick"
	"github.com/GodsBoss/gggg/v2/pkg/game"
)

func TestHandler(t *testing.T) {
	var stateStart game.StateID = 1
	var stateNext game.StateID = 2

	startInit, startInitCalled := started[string]()

	nextInit, nextInitCalled := started[string]()

	tmpl := game.Template[string]{
		States: map[game.StateID]game.State[string]{
			stateStart: {
				InitHandler: startInit,
				ReceiveTickEventHandler: func(_ string, _ tick.Event) game.NextState {
					return game.SwitchState(stateNext)
				},
			},
			stateNext: {
				InitHandler: nextInit,
			},
		},
		InitalStateID: stateStart,
		CreateData: func() string {
			return ""
		},
	}

	instance := tmpl.NewInstance()

	if !*startInitCalled {
		t.Errorf("expected start init handler to be called")
	}
	if *nextInitCalled {
		t.Errorf("did not expect next init handler to be called")
	}

	instance.ReceiveTickEvent(tick.Event{})

	if !*nextInitCalled {
		t.Errorf("expected next init handler to be called")
	}

}

func started[Data any]() (func(Data), *bool) {
	called := false

	return func(_ Data) {
		called = true
	}, &called
}
