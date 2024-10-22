package main

import (
	"fmt"
	"os"
	"errors"
)

const resultsFile = "results.txt"

func main() {
	revenue, err := getUserInput("Revenue: ")
	if (err != nil) {
		panic(err)
	}
	expenses, err := getUserInput("Expenses: ")
	if (err != nil) {
		panic(err)
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if (err != nil) {
		panic(err)
	}

	ebt, eat, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Printf("Earnings: %.0f, Earnings after tax: %.0f, Ratio %.2f\n", ebt, eat, ratio)
	
	writeToFile("Earnings: " + fmt.Sprint(ebt) + "\nEarnings after tax: " + fmt.Sprint(eat) + "\nRatio: " + fmt.Sprint(ratio))
	fmt.Printf("Results have been written to %s", resultsFile)
}

func getUserInput(infoText string) (float64, error){
	fmt.Print(infoText)
	var userInput float64
	
	fmt.Scan(&userInput)
	if (userInput <= 0) {
		return 0, errors.New(infoText + "value is invalid, must be > 0")
	}

	return userInput, nil
}

func calculateFinancials (revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate / 100)
	ratio = ebt / eat
	return
}

func writeToFile(result string) {
	os.WriteFile(resultsFile, []byte(result), 0644)
}