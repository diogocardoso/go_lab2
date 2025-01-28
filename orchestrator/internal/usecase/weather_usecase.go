package usecase

import (
	"errors"

	"github.com.diogocardoso/go/lab-2/orchestrator/internal/entity"
	repository "github.com.diogocardoso/go/lab-2/orchestrator/internal/repositories"
)

type WeatherInput struct {
	Localidade string
	ApiKey     string
}

type GetWeatherUseCase struct {
	WeatherRepository *repository.WeatherRepository
}

func (w *GetWeatherUseCase) Execute(input WeatherInput) (entity.Weather, error) {
	if input.Localidade == "" {
		return entity.Weather{}, errors.New("missing input: Localidade")
	}

	if input.ApiKey == "" {
		return entity.Weather{}, errors.New("missing input: ApiKey")
	}

	weather, err := w.WeatherRepository.Get(input.Localidade, input.ApiKey)
	if err != nil {
		return entity.Weather{}, errors.New("fail to get weather")
	}

	return weather, nil
}
