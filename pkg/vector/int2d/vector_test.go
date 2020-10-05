package int2d_test

import (
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
