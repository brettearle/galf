package storage

import (
	"context"
	"database/sql"
	"fmt"

	fl "github.com/brettearle/galf/internal/flag"
	_ "modernc.org/sqlite"
)

// FOR REFERENCE ONLY
//type Store interface {
//	Create(ctx context.Context, f Flag) error
//	GetByName(ctx context.Context, name string) (Flag, error)
//}

type MemStore struct {
	Store *sql.DB
}

func (m *MemStore) initSchema(ctx context.Context) {
	_, err := m.Store.ExecContext(ctx,
		`CREATE TABLE flag (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255),
			state VARCHAR(255)
		)`)
	if err != nil {
		fmt.Println("failed to create flag table")
	}
}

func (m *MemStore) Create(ctx context.Context, f *fl.Flag) error {
	_, err := m.Store.ExecContext(ctx, `
		INSERT INTO flag (name, state) VALUES (?,?) 
		`, f.Name, f.State)
	if err != nil {
		return err
	}

	return nil
}

func (m *MemStore) GetByName(ctx context.Context, name string) (*fl.Flag, error) {
	var flag fl.Flag
	rows, err := m.Store.QueryContext(ctx, `
		SELECT name, state FROM flag WHERE name=? 
		`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&flag.Name, &flag.State); err != nil {
			return nil, err
		}
	}
	return &flag, nil
}

func (m *MemStore) GetAll(ctx context.Context) (*[]fl.Flag, error) {
	var flagSlice []fl.Flag
	rows, err := m.Store.QueryContext(ctx, `SELECT * FROM flag`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var id int
		var name string
		var state fl.State
		if err := rows.Scan(&id, &name, &state); err != nil {
			return nil, err
		}
		flagSlice = append(flagSlice, fl.Flag{Name: name, State: state})
		count += 1
	}
	return &flagSlice, nil
}

func (m *MemStore) DeleteByName(ctx context.Context, name string) error {
	_, err := m.Store.ExecContext(ctx, `
		DELETE FROM flag WHERE name=? 
		`, name)
	if err != nil {
		return err
	}
	return nil
}

func NewMemStore(ctx context.Context) (*MemStore, error) {
	dsnURI := "file:memdb1?mode=memory&cache=shared"
	db, err := sql.Open("sqlite", dsnURI)
	if err != nil {
		fmt.Println("failed to connect to db")
		return nil, err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	res := &MemStore{
		Store: db,
	}
	res.initSchema(ctx)
	return res, nil
}
