package main

import "fmt"

func main() {
	priceArray := [4]float64{10.99, 9.99, 40.99, 20.0}
	fmt.Println(priceArray)

	// Slices
	featuredPrices := priceArray[1:3]
	featuredPrices[1] = 999.0
	fmt.Println(featuredPrices)
	fmt.Println(priceArray)
	fmt.Println(len(featuredPrices))
	fmt.Println(cap(featuredPrices))

	// Dynamic Lists
	prices := []float64{10.99, 16.5}
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)
	// For removing elements use slices
}