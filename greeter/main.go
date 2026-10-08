package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"greeter/internal/config"
	"greeter/internal/gen"
	"greeter/internal/handlers"
)

func main() {
	cfg := config.Load()

	srv := handlers.NewServer()
	r := chi.NewRouter()

	handler := gen.HandlerWithOptions(gen.NewStrictHandler(srv, nil), gen.ChiServerOptions{
		BaseRouter:       r,
		ErrorHandlerFunc: handlers.WriteJSONError,
	})

	log.Printf("greeter listening on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
