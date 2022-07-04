// Package weather: This is a weather package for forecasting.
package weather

// CurrentCondition current weather condition.
var CurrentCondition string

// CurrentLocation current weather condition.
var CurrentLocation string

// Forecast function produces a forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
