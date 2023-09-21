package vector2d

import "golang.org/x/exp/constraints"

// Sum produces the sum of the given vectors. If given no arguments, the zero value is returned.
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](vectors ...Vector[T]) Vector[T] {
	var v Vector[T]

	for i := range vectors {
		v.x, v.y = v.x+vectors[i].x, v.y+vectors[i].y
	}

	return v
}

// Invert returns a vector with the same length as v, but pointing in the opposite direction.
func Invert[T constraints.Float | constraints.Integer | constraints.Complex](v Vector[T]) Vector[T] {
	return Cartesian(-v.x, -v.y)
}

// Scaled scales v by f.
func Scaled[T constraints.Float | constraints.Integer | constraints.Complex](v Vector[T], f T) Vector[T] {
	return Cartesian(f*v.x, f*v.y)
}
