// Package weather has a tool that describes  
// the weather condition at a specific location.
package weather
// CurrentCondition represents condition.
var CurrentCondition string
//CurrentLocation represents location.
var CurrentLocation string
// Forecast function that returns the weather condition at some location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
