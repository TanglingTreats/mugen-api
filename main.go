package main

import (
	"fmt"
	"net/http"

	"github.com/TanglingTreats/mugen-typer-api/challenges"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	fmt.Println("Starting RESTful service")
	port := ":8080"

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}))

	// Routes
	router.Route("/challenges", challenges.Routes)

	fmt.Printf("Listening at %s\n", port)
	http.ListenAndServe(port, router)
}
