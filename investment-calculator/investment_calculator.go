package main

import (
	"math"
	"fmt"
)

func main() {
	var invAmount, years, expectedReturnRate float64
	const inflationRate = 2.5

	fmt.Print("Investment Amount ")
	fmt.Scan(&invAmount)

	fmt.Print("Years ")
	fmt.Scan(&years)

	fmt.Print("Expected Return ")
	fmt.Scan(&expectedReturnRate)

	futureValue := invAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Printf("Value of %.2f after %.0f years is %v (+%.2f%%))\n", invAmount, years, futureValue, expectedReturnRate)
	fmt.Printf("Adjusted for inflation that is %.2f\n", math.Round(futureRealValue))
}