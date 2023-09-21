package game_test

import (
	"testing"

	"github.com/GodsBoss/gggg/v2/pkg/game"
)

func TestCreatingInstanceFailsWithoutStates(t *testing.T) {
	tmpl := &game.Template[*string]{}

	instance, err := tmpl.NewInstance()
	if instance != nil {
		t.Errorf("expected no instance, got %+v", instance)
	}
	if err == nil {
		t.Errorf("expected an error")
	}
}
