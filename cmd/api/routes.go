package main

import (
	"net/http"

	h "github.com/brettearle/galf/cmd/api/internal/handlers"
	fl "github.com/brettearle/galf/internal/flag"
)

func addRoutes(m *http.ServeMux, f *fl.Service) {
	m.HandleFunc("GET /health", h.Health)
	m.HandleFunc("POST /register", h.Register(f))
	m.HandleFunc("GET /flag/{name}", h.GetByName(f))
}
