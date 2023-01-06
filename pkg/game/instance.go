package game

import (
	"sync"

	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/event/tick"
)

type Instance[Data any] struct {
	states         map[StateID]State[Data]
	data           Data
	currentStateID StateID
	mutex          sync.Mutex
}

func (instance *Instance[Data]) Data() Data {
	return instance.data
}

func (instance *Instance[Data]) ReceiveKeyEvent(event keyboard.Event) {
	defer instance.unlockAfterLock()
	instance.nextState(instance.currentState().ReceiveKeyEvent(instance.data, event))
}

func (instance *Instance[Data]) ReceiveMouseEvent(event mouse.Event) {
	defer instance.unlockAfterLock()
	instance.nextState(instance.currentState().ReceiveMouseEvent(instance.data, event))
}

func (instance *Instance[Data]) ReceiveTickEvent(event tick.Event) {
	defer instance.unlockAfterLock()
	instance.nextState(instance.currentState().ReceiveTickEvent(instance.data, event))
}

func (instance *Instance[Data]) nextState(f func() (StateID, bool)) {
	next, ok := f()
	if ok {
		instance.currentStateID = next
		instance.currentState().Init(instance.data)
	}
}

func (instance *Instance[Data]) unlockAfterLock() func() {
	instance.mutex.Lock()
	return instance.mutex.Unlock
}

func (instance *Instance[Data]) currentState() State[Data] {
	return instance.states[instance.currentStateID]
}
