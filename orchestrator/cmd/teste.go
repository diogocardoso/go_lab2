package main

/*
func teste() {
	// Cria um novo roteador
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Cria uma nova instância do CEPHandler
	cepHandler := handlers.NewCEPHandler()

	// Registra a rota para o CEP
	router.Get("/cep/{cep}", cepHandler.Get)

	// Inicia o servidor em uma goroutine
	go func() {
		log.Println("Starting server on port 8081...")
		if err := http.ListenAndServe(":8081", router); err != nil {
			log.Fatalf("Could not start server: %s\n", err.Error())
		}
	}()

	// Aguarda sinais de encerramento
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down gracefully...")
}
*/
