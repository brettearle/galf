package fl_test

import (
	"testing"

	fl "github.com/brettearle/galf/internal/flag"
	"github.com/brettearle/galf/internal/storage"
)

func TestFlag(t *testing.T) {
	store, err := storage.NewMemStore(t.Context())
	if err != nil {
		t.Fatalf("Couldn't init flag service, aborting test suite before test")
	}
	srv := fl.NewService(store)

	t.Run("Register returns correctly", func(t *testing.T) {
		f := fl.Flag{
			Name:  "feature",
			State: "off",
		}
		err := srv.Register(t.Context(), f)
		if err != nil {
			t.Fatalf("Got %v wanted nil", err)
		}
	})

	t.Run("Register errors with incorrect State value", func(t *testing.T) {
		f := fl.Flag{
			Name:  "feature",
			State: "this is not a state",
		}
		err := srv.Register(t.Context(), f)
		if err == nil {
			t.Fatalf("Got %v wanted error", err)
		}
	})

	t.Run("Register errors with empty name", func(t *testing.T) {
		f := fl.Flag{
			Name:  "",
			State: "off",
		}
		err := srv.Register(t.Context(), f)
		if err == nil {
			t.Fatalf("Got %v wanted error", err)
		}
	})
}
