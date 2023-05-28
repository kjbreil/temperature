package temperature

import "math"

var Round = 1000

func fahrenheitToCelsius(f float64) float64 {
	return round((f - 32) / 1.8)
}

func fahrenheitToKelvin(f float64) float64 {
	return round(fahrenheitToCelsius(f) + 273.15)
}
func celsiusToFahrenheit(c float64) float64 {
	return round(c*1.8 + 32)
}

func celsiusToKelvin(c float64) float64 {
	return round(c + 273.15)
}

func kelvinToCelsius(k float64) float64 {
	return round(k - 273.15)
}

func kelvinToFarenheit(k float64) float64 {
	return round(celsiusToFahrenheit(k - 273.15))
}

func round(input float64) float64 {
	return math.Round(input*float64(Round)) / float64(Round)
}
