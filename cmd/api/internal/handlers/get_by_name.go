package handlers

import (
	"net/http"

	fl "github.com/brettearle/galf/internal/flag"
)

func GetByName(f *fl.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		_, err := f.Get(r.Context(), name)
		if err != nil {
			w.Write([]byte("NOT IMPLEMENTED"))
		}
		//encode flag '_'
		//send down the wire
		w.Write([]byte("NOT IMPLEMENTED"))
	}
}
