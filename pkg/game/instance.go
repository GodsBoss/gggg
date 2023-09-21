package game

import (
	"sync"

	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/event/tick"
)

// Instance represents a game instance. It wraps the data together with a bunch of states and
// takes and passes events.
type Instance[Data any] struct {
	states         map[StateID]State[Data]
	data           Data
	currentStateID StateID
	mutex          sync.Mutex
}

// Data exposes the instance's data.
func (instance *Instance[Data]) Data() Data {
	return instance.data
}

// ReceiveKeyEvent passes a key event to the current state.
func (instance *Instance[Data]) ReceiveKeyEvent(event keyboard.Event) {
	defer instance.unlockAfterLock()()
	instance.nextState(instance.currentState().ReceiveKeyEvent(instance.data, event))
}

// ReceiveMouseEvent passes a mouse event to the current state.
func (instance *Instance[Data]) ReceiveMouseEvent(event mouse.Event) {
	defer instance.unlockAfterLock()()
	instance.nextState(instance.currentState().ReceiveMouseEvent(instance.data, event))
}

// ReceiveTickEvent passes a tick event to the current state.
func (instance *Instance[Data]) ReceiveTickEvent(event tick.Event) {
	defer instance.unlockAfterLock()()
	instance.nextState(instance.currentState().ReceiveTickEvent(instance.data, event))
}

// nextState sets the instance to the next state. If the next state differs from the current state,
// the next state's Init() method is called.
func (instance *Instance[Data]) nextState(f func() (StateID, bool)) {
	next, ok := f()
	if !ok {
		return
	}

	if next == instance.currentStateID {
		return
	}

	instance.currentStateID = next
	instance.currentState().Init(instance.data)
}

// unlockAfterLock locks the instance and returns a function that unlocks it. Use with defer for safe unlocking.
func (instance *Instance[Data]) unlockAfterLock() func() {
	instance.mutex.Lock()
	return instance.mutex.Unlock
}

// currentState returns the current state.
func (instance *Instance[Data]) currentState() State[Data] {
	return instance.states[instance.currentStateID]
}
