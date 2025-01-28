package interfaces

import "github.com.diogocardoso/go/lab-2/orchestrator/internal/entity"

type CEPRepository interface {
	Get(string) ([]byte, error)
	Convert([]byte) (*entity.CEP, error)
	IsValid(string) bool
}
