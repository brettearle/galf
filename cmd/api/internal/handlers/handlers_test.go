package handlers_test

import (
	"testing"

	h "github.com/brettearle/galf/cmd/api/internal/handlers"
)

func Test(t *testing.T) {
	t.Run("Register validate with correct data state 'on'", func(t *testing.T) {
		correct := h.RegisterFlagRequest{Name: "test", State: "on"}
		err := correct.Validate()
		if err != nil {
			t.Errorf("want nil got %v", err)
		}
	})
	t.Run("Register validate with correct data state 'off'", func(t *testing.T) {
		correct := h.RegisterFlagRequest{Name: "test", State: "off"}
		err := correct.Validate()
		if err != nil {
			t.Errorf("want nil got %v", err)
		}
	})

	t.Run("Register errors with empty name", func(t *testing.T) {
		incorrect := h.RegisterFlagRequest{Name: "", State: "on"}
		err := incorrect.Validate()
		if err == nil {
			t.Errorf("want err got nil")
		}
	})
	t.Run("Register errors with incorrect state 'wrong'", func(t *testing.T) {
		incorrect := h.RegisterFlagRequest{Name: "test", State: "wrong"}
		err := incorrect.Validate()
		if err == nil {
			t.Errorf("want err got nil")
		}
	})
}
