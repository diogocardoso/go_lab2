package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com.diogocardoso/go/lab-2/configs"
	"github.com.diogocardoso/go/lab-2/orchestrator/internal/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// Cria um novo roteador
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Cria uma nova instância do CEPHandler
	CEPHandler := handlers.NewCEPHandler(configs)

	router.Get("/cep/{cep}", CEPHandler.Get)

	go func() {
		log.Println("Starting server on port 8081...")
		if err := http.ListenAndServe(":8081", router); err != nil {
			log.Fatalf("Could not start server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down gracefully...")
}
