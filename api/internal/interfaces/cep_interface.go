package interfaces

type CEPInput struct {
	CEP string `json:"cep"`
}

type CEPOutput struct {
	City       string  `json:"city"`
	Celcius    float64 `json:"celcius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Kelvin     float64 `json:"kelvin"`
}
