package handlers

import (
	"encoding/json"
	"net/http"

	fl "github.com/brettearle/galf/internal/flag"
)

func GetAll(f *fl.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flags, err := f.GetAll(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		jsonFlags, err := json.Marshal(flags)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to get flag, please contact admin"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonFlags)
	}
}
