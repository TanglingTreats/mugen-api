package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/TanglingTreats/mugen-api/challenges"
	"github.com/TanglingTreats/mugen-api/dotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	dotenv.InitEnv()

	fmt.Println("Starting RESTful service")
	listenAddr := flag.String("listenaddr", ":8080", "server address")

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

	fmt.Printf("Listening at %s\n", *listenAddr)
	http.ListenAndServe(*listenAddr, router)
}

// Route check for index route
func routeCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// Health check route
// Return 200 OK
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("message: 'OK'"))
}
