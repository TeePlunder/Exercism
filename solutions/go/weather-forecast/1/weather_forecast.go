// Package weather provides the weather for a city.
package weather

// CurrentCondition describes the condition of the city.
var CurrentCondition string
// CurrentLocation is the city where the condition takes place.
var CurrentLocation string
// Forecast gets the condtion for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
