package handlers

import (
	"encoding/json"
	"net/http"

	fl "github.com/brettearle/galf/internal/flag"
)

func GetByName(f *fl.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		flag, err := f.Get(r.Context(), name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		jsonFlag, err := json.Marshal(flag)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to get flag, please contact admin"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonFlag)
	}
}
