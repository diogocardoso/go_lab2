package usecase

import "github.com.diogocardoso/go/lab-2/orchestrator/internal/interfaces"

type ValidateCEPInput struct {
	CEP string
}

type ValidateCEPOutput struct {
	IsValid bool
}

type ValidateCEPUseCase struct {
	CEPRepository interfaces.CEPRepository
}

func NewValidateCEPUseCase(cep interfaces.CEPRepository) *ValidateCEPUseCase {
	return &ValidateCEPUseCase{
		CEPRepository: cep,
	}
}

func (c *ValidateCEPUseCase) Execute(input ValidateCEPInput) bool {
	return c.CEPRepository.IsValid(input.CEP)
}
