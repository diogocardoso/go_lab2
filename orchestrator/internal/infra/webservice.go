package infra

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Port     string
	Handlers map[string]http.HandlerFunc
}

// NewWebServer cria uma nova instância do WebServer
func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Port:     serverPort,
		Handlers: make(map[string]http.HandlerFunc),
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

// Start inicia o servidor web
func (s *WebServer) Start() {
	Router := chi.NewRouter()

	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)

	for path, handler := range s.Handlers {
		log.Printf("path %s | %s ", path, handler)
		Router.Get(path, handler)
	}

	log.Printf("Orchestrator: Starting server on port %s...", s.Port)
	if err := http.ListenAndServe(":"+s.Port, Router); err != nil {
		log.Fatalf("Orchestrator: Could not start server: %s\n", err.Error())
	}
}
