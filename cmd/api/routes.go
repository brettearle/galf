package main

import (
	"net/http"

	h "github.com/brettearle/galf/cmd/api/internal/handlers"
)

func addRoutes(m *http.ServeMux) {
	m.HandleFunc("GET /health", h.Health)
	m.HandleFunc("POST /register", h.Register)
}
