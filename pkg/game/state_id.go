package game

// StateID identifies a state. StateIDs are generated by templates when adding a state.
type StateID interface {
	// stateID is private, so StateID can only be implemented by this package.
	stateID()
}

type stateID struct {
	*int
}

func (id stateID) stateID() {}

// newStateID creates a unique state ID.
func newStateID() StateID {
	var i int

	return stateID{
		&i,
	}
}