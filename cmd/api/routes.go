package main

import (
	"net/http"

	h "github.com/brettearle/galf/cmd/api/internal/handlers"
	fl "github.com/brettearle/galf/internal/flag"
)

func addRoutes(m *http.ServeMux, f *fl.Service) {
	m.HandleFunc("GET /api/health", h.Health)
	m.HandleFunc("POST /api/register", h.Register(f))
	m.HandleFunc("GET /api/flag", h.GetAll(f))
	m.HandleFunc("GET /api/flag/{name}", h.GetByName(f))
}
