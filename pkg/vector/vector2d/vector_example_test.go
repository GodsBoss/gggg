package vector2d_test

import (
	"fmt"

	"github.com/GodsBoss/gggg/v2/pkg/vector/vector2d"
)

func ExampleSum() {
	zero := vector2d.Sum[int]()
	fmt.Printf("(%d, %d)\n", zero.X(), zero.Y())

	sum := vector2d.Sum(vector2d.Cartesian(3, -1), vector2d.Cartesian(-2, 8), vector2d.Cartesian(4, -3))
	fmt.Printf("(%d, %d)\n", sum.X(), sum.Y())

	// Output:
	// (0, 0)
	// (5, 4)
}

func ExampleInvert() {
	inverted := vector2d.Invert(vector2d.Cartesian(8, -5))
	fmt.Printf("(%d, %d)\n", inverted.X(), inverted.Y())

	// Output:
	// (-8, 5)
}

func ExampleScale() {
	scaled := vector2d.Scale(vector2d.Cartesian(3, -4), 5)
	fmt.Printf("(%d, %d)\n", scaled.X(), scaled.Y())

	// Output:
	// (15, -20)
}
