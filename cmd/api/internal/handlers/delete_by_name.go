package handlers

import (
	"net/http"
)

func DeleteByName(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NOT IMPLEMENTED"))
	}
}
