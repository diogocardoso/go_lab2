package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com.diogocardoso/go/lab-2/api/internal/interfaces"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type CEPRepository struct {
	OrchestratorHost string
	OrchestratorPort string
}

func NewCEPRepository(host string, port string) *CEPRepository {
	return &CEPRepository{
		OrchestratorHost: host,
		OrchestratorPort: port,
	}
}

func (r *CEPRepository) IsValid(cep_address string) bool {
	check, _ := regexp.MatchString("^[0-9]{8}$", cep_address)
	return (len(cep_address) == 8 && cep_address != "" && check)
}

func (r *CEPRepository) Get(cep_address string) (*interfaces.CEPOutput, error) {
	log.Println("Context init 1")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	url := fmt.Sprintf("http://%s:%s/cep/%s", r.OrchestratorHost, r.OrchestratorPort, cep_address)

	// Cria a requisição
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Fail to create the request: %v", err)
		return nil, err
	}
	time.Sleep(time.Second * 1)
	// Injeta o contexto no cabeçalho da requisição
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	// Cria o cliente HTTP
	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport,
			otelhttp.WithSpanNameFormatter(func(_ string, req *http.Request) string {
				return "Get-cep-temperature"
			}),
		),
	}

	// Faz a requisição
	resp, err := client.Do(req) // Use req aqui
	if err != nil {
		log.Printf("Fail to make the request: %v", err)
		return nil, err
	}
	defer resp.Body.Close() // Fecha o corpo da resposta

	// Decodifica a resposta JSON
	var cepOutput interfaces.CEPOutput
	if err := json.NewDecoder(resp.Body).Decode(&cepOutput); err != nil {
		log.Printf("Error decoding JSON response: %v", err)
		return nil, err
	}

	return &cepOutput, nil
}
