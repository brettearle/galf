package fl

import (
	"context"
	"fmt"
	"strings"
)

type Store interface {
	Create(ctx context.Context, f *Flag) error
	GetByName(ctx context.Context, name string) (*Flag, error)
	GetAll(ctx context.Context) ([]*Flag, error)
}

type ValidationError struct {
	Fields []string
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

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s", strings.Join(v.Fields, ", "))
}

func (s *Service) Validate(f *Flag) error {
	var failed []string
	if f.Name == "" {
		failed = append(failed, "name is required")
	}

	switch f.State {
	case On, Off:
		//OK
	default:
		failed = append(failed, "state must be: 'on' 'off'")
	}

	if len(failed) > 0 {
		return &ValidationError{Fields: failed}
	}

	return nil
}

func (s *Service) Register(ctx context.Context, f *Flag) error {
	err := s.Validate(f)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	err = s.store.Create(ctx, f)
	if err != nil {
		return fmt.Errorf("Failed to register new flag")
	}
	return nil
}

func (s *Service) Get(ctx context.Context, name string) (*Flag, error) {
	res, err := s.store.GetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("Failed to get flag")
	}
	return res, nil
}
func (s *Service) GetAll(ctx context.Context) ([]*Flag, error) {
	res, err := s.store.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to get flag")
	}
	return res, nil
}
