package entity

import (
	"fmt"
	"strconv"
)

type Weather struct {
	City       string
	Celcius    float64
	Fahrenheit float64
	Kelvin     float64
}

func NewWeather(city string, celcius float64) *Weather {
	return &Weather{
		City:    city,
		Celcius: celcius,
	}
}

func (w *Weather) MakeConversions() {
	w.Fahrenheit, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", w.Celcius*1.8+32), 64)
	w.Kelvin, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", w.Celcius+273.15), 64)
}
