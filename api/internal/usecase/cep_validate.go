package usecase

import (
	"github.com.diogocardoso/go/lab-2/api/internal/repositories"
)

type ValidateCEPInput struct {
	CEP string
}

type ValidateCEPOutput struct {
	IsValid bool
}

type ValidateCEPUseCase struct {
	CEPRepository repositories.CEPRepository
}

func NewValidateCEPUseCase(cep repositories.CEPRepository) *ValidateCEPUseCase {
	return &ValidateCEPUseCase{
		CEPRepository: cep,
	}
}

func (c *ValidateCEPUseCase) Execute(input ValidateCEPInput) bool {
	return c.CEPRepository.IsValid(input.CEP)
}
