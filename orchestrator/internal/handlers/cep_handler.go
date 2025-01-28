package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com.diogocardoso/go/lab-2/configs"
	"github.com.diogocardoso/go/lab-2/orchestrator/internal/interfaces"
	repository "github.com.diogocardoso/go/lab-2/orchestrator/internal/repositories"
	"github.com.diogocardoso/go/lab-2/orchestrator/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type CEPHandler struct {
	CEPRepository     interfaces.CEPRepository
	WeatherRepository *repository.WeatherRepository
	Configs           *configs.Conf
}

func NewCEPHandler(conf *configs.Conf) *CEPHandler {
	return &CEPHandler{
		CEPRepository: repository.NewCEPRepository(),
		Configs:       conf,
	}
}

func NewCEPWeatherHandler(cepRepository interfaces.CEPRepository, weatherRepository *repository.WeatherRepository, configs *configs.Conf) *CEPHandler {
	return &CEPHandler{
		CEPRepository:     cepRepository,
		WeatherRepository: weatherRepository,
		Configs:           configs,
	}
}

func (h *CEPHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Extrai o parâmetro cep da URL
	cep := chi.URLParam(r, "cep")

	// Verifica se o cep está vazio
	if cep == "" {
		http.Error(w, "CEP not found", http.StatusBadRequest)
		return
	}

	getCEP := usecase.NewGetCEPUseCase(h.CEPRepository)
	location, err := getCEP.Location(cep)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting cep: %v", err), http.StatusInternalServerError)
		return
	}

	weather, err := h.WeatherRepository.Get(location, h.Configs.WEATHERMAP_API_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(weather); err != nil {
		log.Printf("Error encoding response to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
