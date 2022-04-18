package keyboard

import "strconv"

// Event represents a keyboard input event.
type Event struct {
	Type EventType

	Location Location

	Alt   bool
	Ctrl  bool
	Shift bool

	Key string
}

type EventType string

const (
	Up    EventType = "up"
	Down  EventType = "down"
	Press EventType = "press"
)

func IsUpEvent(event Event) bool {
	return event.Type == Up
}

func IsDownEvent(event Event) bool {
	return event.Type == Down
}

func IsPressEvent(event Event) bool {
	return event.Type == Press
}

type Location int

const (
	Standard Location = iota
	Left
	Right
	Numpad
)

func IsStandardLocation(event Event) bool {
	return event.Location == Standard
}

func IsLeftLocation(event Event) bool {
	return event.Location == Left
}

func IsRightLocation(event Event) bool {
	return event.Location == Right
}

func IsNumpadLocation(event Event) bool {
	return event.Location == Numpad
}

func (location Location) String() string {
	switch location {
	case Standard:
		return "standard"
	case Left:
		return "left"
	case Right:
		return "right"
	case Numpad:
		return "numpad"
	default:
		return "unknown location " + strconv.Itoa(int(location))
	}
}
