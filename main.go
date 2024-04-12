package main

import (
	"fmt"
	"net/http"

	"github.com/TanglingTreats/mugen-typer-api/challenges"
	"github.com/TanglingTreats/mugen-typer-api/dotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	dotenv.InitEnv()

	fmt.Println("Starting RESTful service")
	port := ":8080"

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	allowedOrigins := []string{"https://mugentyper.com", "https://www.mugentyper.com", "http://localhost:3000"}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// Routes
	router.Get("/", routeCheck)
	router.Get("/health", healthCheck)

	router.Route("/challenges", challenges.Routes)

	fmt.Printf("Listening at %s\n", port)
	http.ListenAndServe(port, router)
}

func routeCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("message: 'OK'"))
}
