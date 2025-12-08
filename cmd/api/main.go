package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	fl "github.com/brettearle/galf/internal/flag"
)

type Config struct {
	Host string
	Port string
}

func NewServer(fl *fl.Service) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, fl)
	var handler http.Handler = mux
	return handler
}

func Run(ctx context.Context, cfg Config, stderr io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx)
	defer cancel()

	//TODO: Place holder WIP
	var store fl.Store
	flagSrv := fl.NewService(store)

	srv := NewServer(flagSrv)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: srv,
	}
	go func() {
		fmt.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil

}

func main() {
	ctx := context.Background()

	err := Run(ctx, Config{
		Host: "0.0.0.0",
		Port: "8080",
	}, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
