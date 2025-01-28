package usecase

import (
	"fmt"

	"github.com.diogocardoso/go/lab-2/api/internal/interfaces"
	"github.com.diogocardoso/go/lab-2/api/internal/repositories"
)

type GetCEPUseCase struct {
	CEPRepository repositories.CEPRepository
}

func NewGetCEPUseCase(CEPRepository repositories.CEPRepository) *GetCEPUseCase {
	return &GetCEPUseCase{
		CEPRepository: CEPRepository,
	}
}

func (c *GetCEPUseCase) Execute(cep string) (interfaces.CEPOutput, error) {
	cepOutput, err := c.CEPRepository.Get(cep)
	if err != nil {
		return interfaces.CEPOutput{}, err
	}
	if cepOutput == nil {
		return interfaces.CEPOutput{}, fmt.Errorf("cep not found")
	}

	return *cepOutput, nil
}
