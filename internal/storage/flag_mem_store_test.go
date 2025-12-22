package storage_test

import (
	"context"
	"testing"

	fl "github.com/brettearle/galf/internal/flag"
	"github.com/brettearle/galf/internal/storage"
)

func TestMemStore(t *testing.T) {
	ctx := t.Context()
	db, err := storage.NewMemStore(ctx)
	if err != nil {
		t.Errorf("failed to init DB")
	}
	t.Cleanup(func() { db.Store.Close() })

	t.Run("ping store", func(t *testing.T) {
		err = db.Store.Ping()
		if err != nil {
			t.Errorf("failed to ping db")
		}
	})

	t.Run("schema should include flag table with id, name, state columns", func(t *testing.T) {
		var tableName string
		err := db.Store.QueryRow(`
		SELECT name 
		FROM sqlite_master 
		WHERE type = 'table' AND name = 'flag'`).Scan(&tableName)
		if err != nil {
			t.Fatalf("expected table %q to exist, but it does not: %v", "flag", err)
		}

		rows, err := db.Store.Query(`PRAGMA table_info(flag)`)
		if err != nil {
			t.Fatalf("failed to query table info for flag: %v", err)
		}
		defer rows.Close()

		cols := map[string]bool{}
		for rows.Next() {
			var (
				cid       int
				name      string
				colType   string
				notNull   int
				dfltValue any
				pk        int
			)
			if err := rows.Scan(&cid, &name, &colType, &notNull, &dfltValue, &pk); err != nil {
				t.Fatalf("failed to scan table_info row: %v", err)
			}
			cols[name] = true
		}
		if err := rows.Err(); err != nil {
			t.Fatalf("rows error: %v", err)
		}

		wantCols := []string{"id", "name", "state"}
		for _, c := range wantCols {
			if !cols[c] {
				t.Errorf("expected column %q on table flag, but it was missing", c)
			}
		}
	})

	t.Run("Create a flag with name `feature` and state `off`", func(t *testing.T) {
		flag := fl.Flag{
			Name:  "feature",
			State: "off",
		}
		err := db.Create(ctx, &flag)
		if err != nil {
			t.Errorf("failed to create flag row %v", err)
		}
	})

	t.Run("Get flag with name `feature`", func(t *testing.T) {
		flag := fl.Flag{
			Name:  "feature",
			State: "off",
		}
		err := db.Create(ctx, &flag)
		if err != nil {
			t.Fatalf(".Create(ctx, %v) got error %v want nil", flag, err)
		}

		got, err := db.GetByName(ctx, flag.Name)
		if err != nil {
			t.Fatalf("GetByName(ctx, %s) got error %v want nil", flag.Name, err)
		}

		if got.Name != flag.Name {
			t.Fatalf(".GetByName(ctx, %s) got %s want %s", flag.Name, got.Name, flag.Name)
		}
	})

	t.Run("Get all flags", func(t *testing.T) {
		db.Store.Exec(`DELETE FROM flag`)

		flag1 := fl.Flag{
			Name:  "feature1",
			State: "off",
		}
		flag2 := fl.Flag{
			Name:  "feature2",
			State: "off",
		}
		err = db.Create(ctx, &flag1)
		if err != nil {
			t.Fatalf(".Create(ctx, %v) got error %v want nil", flag1, err)
		}
		err = db.Create(ctx, &flag2)
		if err != nil {
			t.Fatalf(".Create(ctx, %v) got error %v want nil", flag2, err)
		}

		got, err := db.GetAll(ctx)
		if err != nil {
			t.Fatalf("GetAll(ctx) got error %v want nil", err)
		}
		deref := *got
		if deref[0].Name != flag1.Name {
			t.Fatalf(".GetAll(ctx) f1 got %s want %s", deref[0].Name, flag1.Name)
		}
		if deref[1].Name != flag2.Name {
			t.Fatalf(".GetAll(ctx) f2 got %s want %s", deref[1].Name, flag2.Name)
		}
	})
}

func TestMemStore_DeleteByName(t *testing.T) {
	tests := []struct {
		id string // description of this test case
		// Named input parameters for target function.
		name    string
		wantErr bool
		//Setup required
		flags []fl.Flag
	}{
		{
			id:      "delete success",
			name:    "featureToDelete",
			wantErr: false,
			flags: []fl.Flag{
				{
					Name:  "featureToDelete",
					State: "off",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			m, err := storage.NewMemStore(context.Background())
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			for _, f := range tt.flags {
				err = m.Create(t.Context(), &f)
				if err != nil {
					t.Fatalf(".Create(ctx, %v) got error %v want nil", f, err)
				}
			}

			gotErr := m.DeleteByName(t.Context(), tt.name)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("DeleteByName() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("DeleteByName() succeeded unexpectedly")
			}
		})
	}
}
