// Package weather is used to forecast weather of a location.
package weather

// CurrentCondition holds weather condition.
var CurrentCondition string

// CurrentLocation holds city.
var CurrentLocation string

// Forecast returns weather forcase for given city using given condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
