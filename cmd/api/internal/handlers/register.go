package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	fl "github.com/brettearle/galf/internal/flag"
)

type RegisterFlagRequest struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type ValidationError struct {
	Fields []string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s", strings.Join(v.Fields, ", "))
}

func (f *RegisterFlagRequest) Validate() error {
	var failed []string
	if f.Name == "" {
		failed = append(failed, "name is required")
	}

	switch f.State {
	case string(fl.On), string(fl.Off):
		//OK
	default:
		failed = append(failed, "state must be: 'on' 'off'")
	}

	if len(failed) > 0 {
		return &ValidationError{Fields: failed}
	}

	return nil
}

func (f *RegisterFlagRequest) ToFlag() fl.Flag {
	return fl.Flag{
		Name:  f.Name,
		State: fl.State(f.State),
	}
}

func Register(f *fl.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var v RegisterFlagRequest

		ct := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ct, "application/json") {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := v.Validate(); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		flag := v.ToFlag()

		err := f.Register(r.Context(), flag)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
