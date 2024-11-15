package main

import (
	"fmt"
	"log"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for index, taxRate := range(taxRates) {
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)

		//cmdm := cmdmanager.New()
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		pricesJob := prices.NewTaxIncludedPriceJob(*fm, taxRate)
		go pricesJob.Process(doneChans[index], errChans[index])
	}

	for index := range taxRates {
		select {
		case err := <- errChans[index]:
			if err != nil {
				log.Fatalf("Error processing the job: %v", err)
			}
		case <- doneChans[index]:
			fmt.Println("Job done")
		}
	}
}