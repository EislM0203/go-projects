package main

import "fmt"

func main() {
	age := 32
	fmt.Println("Age: ", age)
	
	adultYears := getAdultYears(&age)
	fmt.Println("Adult years:", adultYears)

	setAgeToAdultYears(&age)
	fmt.Println("Adult years (age): ", age)
}

func getAdultYears(age *int) int {
	return *age - 18
}

func setAgeToAdultYears(age *int) {
	*age = *age - 18
}