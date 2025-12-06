package main

import (
	"net/http"

	fl "github.com/brettearle/galf/cmd/api/internal/flag"
	h "github.com/brettearle/galf/cmd/api/internal/handlers"
)

func addRoutes(m *http.ServeMux, f *fl.Service) {
	m.HandleFunc("GET /health", h.Health)
	m.HandleFunc("POST /register", h.Register(f))
}
