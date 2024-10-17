package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, eat, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Printf("Earnings: %.0f, Earnings after tax: %.0f, Ratio %.2f", ebt, eat, ratio)
}

func getUserInput(infoText string) (userInput float64){
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials (revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate / 100)
	ratio = ebt / eat
	return
}