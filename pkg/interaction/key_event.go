package interaction

type KeyEvent struct {
	Type KeyEventType

	Alt   bool
	Ctrl  bool
	Shift bool

	Key string
}

type KeyEventType string

const (
	KeyUp    KeyEventType = "up"
	KeyDown  KeyEventType = "down"
	KeyPress KeyEventType = "press"
)
