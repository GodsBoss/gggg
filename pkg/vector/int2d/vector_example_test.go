package int2d_test

import (
	"fmt"

	"github.com/GodsBoss/gggg/pkg/vector/int2d"
)

func ExampleFromXY() {
	v := int2d.FromXY(4, 7)
	fmt.Printf("(%d, %d)\n", v.X(), v.Y())

	// Output:
	// (4, 7)
}

func ExampleAdd() {
	v1 := int2d.FromXY(-2, 6)
	v2 := int2d.FromXY(4, 3)
	v := int2d.Add(v1, v2)
	fmt.Printf("(%d, %d)\n", v.X(), v.Y())

	// Output:
	// (2, 9)
}

func ExampleMultiply() {
	v := int2d.Multiply(int2d.FromXY(3, -7), 4)
	fmt.Printf("(%d, %d)\n", v.X(), v.Y())

	// Output:
	// (12, -28)
}
