package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com.diogocardoso/go/lab-2/api/internal/interfaces"
	"github.com.diogocardoso/go/lab-2/api/internal/repositories"
	"github.com.diogocardoso/go/lab-2/api/internal/usecase"
	"github.com.diogocardoso/go/lab-2/configs"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type CEPHandler struct {
	CEPRepository repositories.CEPRepository
	Configs       *configs.Conf
	Tracer        trace.Tracer
}

func NewCEPHandler(conf *configs.Conf, tracer trace.Tracer) *CEPHandler {
	return &CEPHandler{
		CEPRepository: *repositories.NewCEPRepository(conf.ORCHESTRATOR_HOST, conf.ORCHESTRATOR_PORT),
		Tracer:        tracer,
	}
}

func (h *CEPHandler) Get(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, span := h.Tracer.Start(ctx, "Validate-cep")
	defer span.End()

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("fail to read the response: %v", err), http.StatusInternalServerError)
		return
	}

	var cep_data usecase.ValidateCEPInput
	err = json.Unmarshal(resp, &cep_data)
	if err != nil {
		http.Error(w, fmt.Sprintf("fail to parse the cep_data: %v", err), http.StatusInternalServerError)
		return
	}

	validate_cep_dto := usecase.ValidateCEPInput{
		CEP: cep_data.CEP,
	}

	validateCEP := usecase.NewValidateCEPUseCase(h.CEPRepository)
	is_valid := validateCEP.Execute(validate_cep_dto)
	if !is_valid {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	get_cep_dto := interfaces.CEPInput{
		CEP: cep_data.CEP,
	}

	getCEP := usecase.NewGetCEPUseCase(h.CEPRepository)
	Output, err := getCEP.Execute(get_cep_dto.CEP)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Codifica o mapa em JSON e escreve na resposta
	if err := json.NewEncoder(w).Encode(Output); err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
	}

	log.Printf("|---> Success: %v | C: %v | K: %v | F:%v", Output.City, Output.Celcius, Output.Kelvin, Output.Fahrenheit)
}
