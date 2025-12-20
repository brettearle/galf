package main

import (
	"net/http"

	h "github.com/brettearle/galf/cmd/api/internal/handlers"
	fl "github.com/brettearle/galf/internal/flag"
)

func addRoutes(m *http.ServeMux, f *fl.Service) {
	m.HandleFunc("GET /api/health", h.Health)
	m.HandleFunc("POST /api/register", h.Register(f))
	m.HandleFunc("GET /api/flags", h.GetAll(f))
	m.HandleFunc("GET /api/flag/{name}", h.GetByName(f))
	//TODO: Delete flag by name
	m.HandleFunc("DELETE /api/flag/{name}", h.DeleteByName(f))
	//TODO: Update flag by name
	//TODO: Update state by name
}
