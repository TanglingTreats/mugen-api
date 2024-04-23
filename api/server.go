package api

import (
	"net/http"
	"strings"

	"github.com/TanglingTreats/mugen-api/challenges"
	"github.com/TanglingTreats/mugen-api/dotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	router     chi.Router
	listenAddr string
}

func NewServer(listenAddr string) Server {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	envOrigins := dotenv.GetEnvVar("ALLOWED_ORIGIN")
	allowedOrigins := strings.Split(envOrigins, " ")

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	return Server{router, listenAddr}
}

func (s *Server) Start() error {
	// Routes
	s.router.Get("/", index)
	s.router.Get("/health", healthCheck)

	s.router.Route("/challenges", challenges.Routes)

	return http.ListenAndServe(s.listenAddr, s.router)
}

// index route
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("message: 'OK'"))
}

// Health check route
// Return 200 OK
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("message: 'OK'"))
}
