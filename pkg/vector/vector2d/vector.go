package vector2d

func Cartesian[T any](x T, y T) Vector[T] {
	return Vector[T]{
		x: x,
		y: y,
	}
}

// Vector is a value type representing a 2D vector. The zero value is the vector (0, 0).
type Vector[T any] struct {
	x T
	y T
}

func (v Vector[T]) X() T {
	return v.x
}

func (v Vector[T]) Y() T {
	return v.y
}

// Zero returns a zero-value vector.
func Zero[T any]() Vector[T] {
	return Vector[T]{}
}
