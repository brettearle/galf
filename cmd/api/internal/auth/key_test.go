package auth_test

import (
	"testing"

	"github.com/brettearle/galf/cmd/api/internal/auth"
)

func TestValidatePrefix(t *testing.T) {
	t.Run("Valid status", func(t *testing.T) {
		status := "live"
		got := auth.ValidatePrefix(status)
		if got != true {
			t.Errorf("invalid status %v", status)
		}
	})

	t.Run("Valid status", func(t *testing.T) {
		status := "test"
		got := auth.ValidatePrefix(status)
		if got != true {
			t.Errorf("invalid status %v", status)
		}
	})

	t.Run("Invalid status", func(t *testing.T) {
		status := "not a status"
		got := auth.ValidatePrefix(status)
		if got != false {
			t.Errorf("invalid status %v", status)
		}
	})
}
