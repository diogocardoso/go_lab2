package usecase

import (
	"strings"

	"github.com.diogocardoso/go/lab-2/orchestrator/internal/interfaces"
)

type CEPInput struct {
	CEP string
}

type CEPOutput struct {
	Localidade string `json:"localidade"`
}

type GetCEPUseCase struct {
	CEPRepository interfaces.CEPRepository
}

func NewGetCEPUseCase(cep interfaces.CEPRepository) *GetCEPUseCase {
	return &GetCEPUseCase{
		CEPRepository: cep,
	}
}

func (c *GetCEPUseCase) Location(cep string) (string, error) {
	cep_resp, err := c.CEPRepository.Get(cep)
	if err != nil && strings.Contains(string(cep_resp), "Http 400") {
		return "", err
	}

	cep_, err := c.CEPRepository.Convert(cep_resp)
	if err != nil {
		return "", err
	}

	return cep_.Localidade, nil
}
