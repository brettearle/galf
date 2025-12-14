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
			w.Write([]byte("NOT IMPLEMENTED"))
		}

		_, err = json.Marshal(flag)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to get flag, please contact admin"))
		}
		//TODO:
		//encode flag '_' this encode should be its own package so we dont have to re write
		//send down the wire

		w.Write([]byte("NOT IMPLEMENTED"))
	}
}
