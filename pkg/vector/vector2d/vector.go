package vector2d

import "golang.org/x/exp/constraints"

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

type Numeric interface {
	constraints.Float | constraints.Integer
}

// Sum produces the sum of the given vectors. If given no arguments, the zero value is returned.
func Sum[T Numeric](vectors ...Vector[T]) Vector[T] {
	var v Vector[T]

	for i := range vectors {
		v.x, v.y = v.x+vectors[i].x, v.y+vectors[i].y
	}

	return v
}

// Invert returns a vector with the same length as v, but pointing in the opposite direction.
func Invert[T Numeric](v Vector[T]) Vector[T] {
	return Cartesian(-v.x, -v.y)
}

// Scaled scales v by f.
func Scaled[T Numeric](v Vector[T], f T) Vector[T] {
	return Cartesian(f*v.x, f*v.y)
}
