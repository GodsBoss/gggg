package mouse

type Event struct {
	Type EventType

	Button Button

	Alt   bool
	Ctrl  bool
	Shift bool

	X int
	Y int
}

type Button int

const (
	Primary Button = iota
	Auxiliary
	Secondary
)

// IsPrimaryButtonEvent checks wether the primary button (usually left) was the cause of an up or down event.
func IsPrimaryButtonEvent(event Event) bool {
	return IsButtonEvent(event) && event.Button == 0
}

// IsSecondaryButtonEvent checks wether the secondary button (usually right) was the cause of an up or down event.
func IsSecondaryButtonEvent(event Event) bool {
	return IsButtonEvent(event) && event.Button == 2
}

// IsAuxiliaryButtonEvent checks wether the auxiliary button (usually middle / mouse wheel) was the cause of an up or down event.
func IsAuxiliaryButtonEvent(event Event) bool {
	return IsButtonEvent(event) && event.Button == 1
}

type EventType string

const (
	Up   EventType = "up"
	Down EventType = "down"
	Move EventType = "move"
)

func IsButtonEvent(event Event) bool {
	return IsUpEvent(event) || IsDownEvent(event)
}

func IsUpEvent(event Event) bool {
	return event.Type == Up
}

func IsDownEvent(event Event) bool {
	return event.Type == Down
}

func IsMoveEvent(event Event) bool {
	return event.Type == Move
}
