package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com.diogocardoso/go/lab-2/orchestrator/internal/entity"
)

type CEPRepository struct{}

func NewCEPRepository() *CEPRepository {
	return &CEPRepository{}
}

func (r *CEPRepository) IsValid(cep string) bool {
	check, _ := regexp.MatchString("^[0-9]{8}$", cep)
	return (len(cep) == 8 && cep != "" && check)
}

func (r *CEPRepository) Get(cep_address string) ([]byte, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep_address)

	res, err := http.Get(url)
	if err != nil {
		log.Printf("Erro ao fazer requisição para o CEP %s: %s", cep_address, err)
		return nil, err
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Fail to read the response: %v", err)
		return nil, err
	}

	return resp, nil
}

func (r *CEPRepository) Convert(cep_response []byte) (*entity.CEP, error) {
	var cep entity.CEP

	err := json.Unmarshal(cep_response, &cep)
	if err != nil {
		log.Printf("Fail to decode the response: %v", err)
		return nil, err
	}
	return &cep, nil
}
