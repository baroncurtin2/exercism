// Package weather provides the Forecast for the weather.
package weather

// CurrentCondition: stores the current condition of goblinocus.
var CurrentCondition string
// CurrentLocation: stores the current location of thegoblinocus.
var CurrentLocation string

// Forecast of the CurrentLocation with the CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
