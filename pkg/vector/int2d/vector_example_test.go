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
