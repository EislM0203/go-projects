package main

import (
	"fmt"
	"log"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range(taxRates) {
		//cmdm := cmdmanager.New()
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		pricesJob := prices.NewTaxIncludedPriceJob(*fm, taxRate)
		err := pricesJob.Process()

		if err != nil {
			log.Fatalf("Error processing the job: %v", err)
		}
	}
}