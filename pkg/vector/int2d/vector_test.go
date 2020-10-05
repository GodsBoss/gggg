package int2d_test

import (
	"fmt"
	"testing"

	"github.com/GodsBoss/gggg/pkg/vector/int2d"
)

func TestZeroVector(t *testing.T) {
	z := int2d.Zero()

	if z.X() != 0 {
		t.Errorf("expected zero.X() to be 0, not %d", z.X())
	}
	if z.Y() != 0 {
		t.Errorf("expected zero.Y() to be 0, not %d", z.Y())
	}
}

func TestRotateClockwise(t *testing.T) {
	testCases := map[int2d.Vector]int2d.Vector{
		int2d.Up():    int2d.Right(),
		int2d.Right(): int2d.Down(),
		int2d.Down():  int2d.Left(),
		int2d.Left():  int2d.Up(),
	}

	for input, expectedResult := range testCases {
		t.Run(
			fmt.Sprintf("(%d, %d)", input.X(), input.Y()),
			func(t *testing.T) {
				actualResult := int2d.RotateClockwise(input)
				if actualResult != expectedResult {
					t.Errorf(
						"expected RotateClockwise((%d, %d)) to be (%d, %d), but got (%d, %d)",
						input.X(), input.Y(),
						expectedResult.X(), expectedResult.Y(),
						actualResult.X(), actualResult.Y(),
					)
				}
			},
		)
	}
}
