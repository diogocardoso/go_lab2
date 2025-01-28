package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com.diogocardoso/go/lab-2/api/internal/handlers"
	"github.com.diogocardoso/go/lab-2/api/internal/infra"
	"github.com.diogocardoso/go/lab-2/collector"
	"github.com.diogocardoso/go/lab-2/configs"
	orchestrator "github.com.diogocardoso/go/lab-2/orchestrator/server"
	"go.opentelemetry.io/otel"
)

func ConfigureServer(conf *configs.Conf) *infra.WebServer {
	tracer := otel.Tracer("input-api-tracer")
	webServer := infra.NewWebServer(conf.API_PORT)

	CEPHandler := handlers.NewCEPHandler(conf, tracer)

	webServer.AddHandler("/cep", CEPHandler.Get)
	return webServer
}

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := collector.InitProvider(configs.APP_NAME, configs.COLLECTOR_HOST)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	go func() {
		webserver := ConfigureServer(configs)
		webserver.Start()
	}()

	go func() {
		err := startOrchestrator()
		if err != nil {
			log.Fatalf("Erro ao iniciar o Orchestrator: %v", err)
		}
	}()

	select {
	case <-sigCh:
		log.Println("Shutting down gracefully, CTRL+c pressed...")
	case <-ctx.Done():
		log.Println("Shutting down due other reason...")
	}

	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
}

func startOrchestrator() error {
	orchestrator.Server()
	return nil
}
