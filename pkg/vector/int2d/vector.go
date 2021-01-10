// Package int2d provides integer 2D vectors, mainly useful for full pixel
// coordinates or rectangular grid coordinates.
//
// The coordinate system represented by this package is as usually found in
// computer graphics: Coordinates are growing when going to the right or down.
package int2d

// Vector represents an integer 2D vector. Vectors are immutable, i.e. you
// can new ones (sometimes from existing ones).
// Vectors can be used safely as map keys.
type Vector struct {
	x int
	y int
}

// X provides the horizontal component of a vector.
func (v Vector) X() int {
	return v.x
}

// Y provides the  vertical component of a vector.
func (v Vector) Y() int {
	return v.y
}

// Zero returns the zero vector (0, 0).
func Zero() Vector {
	return Vector{}
}

// FromXY creates a new vector (x, y).
func FromXY(x int, y int) Vector {
	return Vector{
		x: x,
		y: y,
	}
}

// Up returns the smallest vector pointing upwards, (0, -1).
func Up() Vector {
	return FromXY(0, -1)
}

// Right returns the smallest vector pointing right, (1, 0).
func Right() Vector {
	return FromXY(1, 0)
}

// Down returns the smallest vector pointing downwards, (0, 1).
func Down() Vector {
	return FromXY(0, 1)
}

// Left returns the smallest vector pointing left, (-1, 0).
func Left() Vector {
	return FromXY(-1, 0)
}

// RotateClockwise takes a vector and rotates it clockwise.
func RotateClockwise(v Vector) Vector {
	return FromXY(-v.Y(), v.X())
}

// RotateCounterclockwise takes a vector and rotates it counterclockwise.
func RotateCounterclockwise(v Vector) Vector {
	return FromXY(v.Y(), -v.X())
}

// Add creates a new vector which is the sum of all the vectors passed to this
// function. If no vector is given, the zero vector is returned.
func Add(vectors ...Vector) Vector {
	v := Vector{}
	for i := range vectors {
		v.x += vectors[i].x
		v.y += vectors[i].y
	}
	return v
}

// Multiply creates a new vector by multiplying both horizontal and vertical
// component of v with f.
func Multiply(v Vector, f int) Vector {
	return Vector{
		x: v.x * f,
		y: v.y * f,
	}
}

// Divide creates a new vector by dividing both horizontal and vertical
// components of v by d. Rounds if necessary.
func Divide(v Vector, d int) Vector {
	return Vector{
		x: v.x / d,
		y: v.y / d,
	}
}
