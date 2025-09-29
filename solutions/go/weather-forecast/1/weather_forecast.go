// Package weather provides forecast of given city.
package weather

// CurrentCondition represents a cetain condition of thr given city.
var CurrentCondition string
// CurrentLocation represents a cetain location.
var CurrentLocation string

// Forecast returns an string value stating CurrentLocation and CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
