package main

import "fmt"

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	switch speed {
	case 1, 2, 3, 4:
		return 1.0
	case 5, 6, 7, 8:
		return 0.9
	case 9, 10:
		return 0.77
	}
	return 0.0
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(speed*221) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

func main() {
	rate := CalculateProductionRatePerMinute(5)
	fmt.Println(rate)

	rate2 := CalculateProductionRatePerHour(7)
	fmt.Println(rate2)
}
