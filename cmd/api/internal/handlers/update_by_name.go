package handlers

import (
	"fmt"
	"net/http"

	fl "github.com/brettearle/galf/internal/flag"
)

func UpdateByName(f *fl.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		err := f.Update(r.Context(), name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		res := fmt.Sprintf("flag: %v deleted", name)
		w.Write([]byte(res))
	}
}
