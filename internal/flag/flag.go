package fl

import (
	"context"
	"fmt"
)

type Store interface {
	Create(ctx context.Context, f Flag) error
	GetByName(ctx context.Context, name string) (*Flag, error)
}

// Chosen incase additional states are required

// State represents the possible on/off states for a flag.
//
// Currently valid values are:
//   - "on"
//   - "off"
type State string

const (
	On  State = "on"
	Off State = "off"
)

type Flag struct {
	Name  string
	State State
}

type Service struct {
	store Store
}

func NewService(s Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) Register(f Flag) {
	fmt.Println("Registering NOT IMPLEMENTED arg = ", f)
}
