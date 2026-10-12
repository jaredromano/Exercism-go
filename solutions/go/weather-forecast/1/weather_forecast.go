// Package weather provides functionality to track and forecast the current weather conditions for a given location.
package weather

var (
	// CurrentCondition holds the current weather condition for the tracked location.
	CurrentCondition string
	// CurrentLocation holds the current location for which the weather is being tracked.
	CurrentLocation string
)

// Forecast updates the current weather condition for the specified city and returns a formatted string with the current location and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
