package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com.diogocardoso/go/lab-2/orchestrator/internal/entity"
)

type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return e.Message
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WeatherRepository struct {
	client HTTPClient
}

func NewWeatherRepository(client HTTPClient) *WeatherRepository {
	return &WeatherRepository{
		client: client,
	}
}

func (w *WeatherRepository) Get(city string, api_key string) (entity.Weather, error) {
	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", api_key, encodedCity)
	resp, err := http.Get(url)
	if err != nil {
		return entity.Weather{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Erro URL: %s", url)
		return entity.Weather{}, &HTTPError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("error getting city (%s) weather data", city)}
	}

	var result struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Erro ao decodificar resposta para a localização %s: %s", city, err)
		return entity.Weather{}, err
	}

	weather := entity.NewWeather(city, result.Current.TempC)

	weather.MakeConversions()

	log.Printf("|---> %s | C: %.2f | F: %.2f | K: %.2f", encodedCity, weather.Celcius, weather.Fahrenheit, weather.Kelvin)

	return *weather, nil
}
