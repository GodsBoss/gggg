package game

// Template is a template for game instances.
type Template[Data any] struct {
	States        map[StateID]State[Data]
	InitalStateID StateID
	CreateData    func() Data
}

// NewInstance creates a new game instance from a game template.
func (tmpl Template[Data]) NewInstance() *Instance[Data] {
	var data Data
	if tmpl.CreateData != nil {
		data = tmpl.CreateData()
	}

	instance := &Instance[Data]{
		states: tmpl.States,
		data:   data,
	}
	instance.nextState(SwitchState(tmpl.InitalStateID))

	return instance
}
