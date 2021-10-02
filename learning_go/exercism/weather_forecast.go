package main

// CurrentCondition contains the current weather forecast
var CurrentCondition string

// CurrentLocation contains the current location used by the weather forecast
var CurrentLocation string

// Forecast takes a city & condition that are strings and returns a string of
// the current weather and location
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

func main() {

}
