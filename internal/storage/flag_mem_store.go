package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

//type Store interface {
//	Create(ctx context.Context, f Flag) error
//	GetByID(ctx context.Context, name string) (Flag, error)
//}

type MemStore struct {
	Store *sql.DB
}

func (m *MemStore) initSchema(ctx context.Context) {
	_, err := m.Store.ExecContext(ctx,
		`CREATE TABLE flag (
			id INT PRIMARY KEY,
			name VARCHAR(255),
			state VARCHAR(255)
		)`)
	if err != nil {
		fmt.Println("failed to create flag table")
	}
}

func NewMemStore() (*MemStore, error) {
	ctx := context.Background()
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
