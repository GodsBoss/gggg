package game_test

import (
	"fmt"

	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/event/tick"
	"github.com/GodsBoss/gggg/v2/pkg/game"
)

func ExampleTemplate() {
	template := &game.Template[*string]{
		CreateData: func() *string {
			var s string
			return &s
		},
	}

	// titleScreen represents the game's title screen. As this is the first state added, it becomes the start state.
	titleScreen := template.AddState()

	// start state initializes the game's data. We want this to be the start state.
	start := template.AddState().SetAsStartState()

	// playing is the play state, where the player (hopefully) spends most of their time.
	playing := template.AddState()

	// start will only initialize the game data and then immediately switch to the next state.
	start.SetInitHandler(
		func(data *string) game.NextState {
			*data = "START"
			return game.SwitchState(titleScreen.ID())
		},
	)

	// titleScreen waits for a mouse event, then switches to playing state. In a real game, this would probably
	// only let specific mouse events start playing.
	titleScreen.SetMouseHandler(
		func(data *string, _ mouse.Event) game.NextState {
			fmt.Println("MOUSE EVENT")
			return game.SwitchState(playing.ID())
		},
	)

	// playing sets the game data to the key pressed or released. Also changes the game data on tick events.
	playing.
		SetKeyboardHandler(
			func(data *string, ev keyboard.Event) game.NextState {
				*data = ev.Key
				return game.SameState()
			},
		).SetTickHandler(
		func(data *string, ev tick.Event) game.NextState {
			*data = "TICK"
			return game.SameState()
		},
	)

	// Create a fresh game instance from the template.
	instance, err := template.NewInstance()
	if err != nil {
		fmt.Printf("error: %+v", err)
		return
	}

	// Show current game data.
	fmt.Println(*instance.Data())

	// Send a tick event. Does nothing.
	instance.ReceiveTickEvent(tick.Event{})

	// Send a key event. Does nothing.
	instance.ReceiveKeyEvent(keyboard.Event{})

	// Send a mouse event. State switches from title to playing.
	instance.ReceiveMouseEvent(mouse.Event{})

	// Send a key event. Changes game data to the key.
	instance.ReceiveKeyEvent(
		keyboard.Event{
			Key: "X",
		},
	)

	fmt.Println(*instance.Data())

	// Send a mouse event. Does nothing.
	instance.ReceiveMouseEvent(mouse.Event{})

	// Send a tick event.
	instance.ReceiveTickEvent(tick.Event{})

	fmt.Println(*instance.Data())

	// Output:
	// START
	// MOUSE EVENT
	// X
	// TICK
}
