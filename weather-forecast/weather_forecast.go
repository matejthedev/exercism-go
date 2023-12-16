// Package weather is used for a weather forecast.
package weather

// CurrentCondition represents a current weather condition.
var CurrentCondition string

// CurrentLocation represents a city.
var CurrentLocation string

// Forecast returns a string explaining a current weather condition in a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
