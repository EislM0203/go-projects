package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")

	var accountBalance, err = readFromFile()
	if (err != nil) {
		fmt.Println("No balance found, creating new checking account!")
		writeToFile(0)
	}
	
	for {
		fmt.Println("What do you want to do?")
		choice := getUserInput("1. Check Balance\n2. Deposit Money\n3. Withdraw Money\n4. Exit\n")
		switch choice {
			case 1:
				fmt.Printf("Your account balance: %.2f€\n", accountBalance)
			case 2:
				deposit := getUserInput("How much money would you like to deposit?\n")
				if deposit <= 0 {
					fmt.Printf("%f is invalid. Must be greater than 0.\n", deposit)
					continue
				}
				accountBalance += deposit
				writeToFile(accountBalance)
				fmt.Printf("New Amount: %f", accountBalance)
			case 3:
				withdrawl := getUserInput("How much money would you like to withdraw?\n")
				if withdrawl < 0 {
					fmt.Printf("%f is invalid. Must be greater than 0.\n", withdrawl)
					continue
				}
				if withdrawl > accountBalance {
					fmt.Printf("You cannot withdraw more than you have. Current balance: %.2f\n", accountBalance)
					continue
				}

				accountBalance -= withdrawl
				writeToFile(accountBalance)
				fmt.Printf("New Amount: %.2f€\n", accountBalance)
			default:
				fmt.Println("See ya! Thanks for choosing our bank!")
				return
		}
	}
}

func getUserInput(infoText string) (userInput float64){
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func writeToFile(balance float64) {
	os.WriteFile(accountBalanceFile, []byte(fmt.Sprint(balance)), 0644)
}

func readFromFile() (float64, error){
	data, err := os.ReadFile(accountBalanceFile)
	if (err != nil) {
		return 0, errors.New("failed to find balance file")
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	if (err != nil) {
		return 0, errors.New("file does not contain a number")
	}
	return balance, nil
}