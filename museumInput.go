// Created by: haokai
// Created on: May 2021
// This program displays, "Age-and-Movie-and-Day"

package main

import (
	"fmt"
)

func main() {

	// This function does addition
	var age int
	var day string

	// input
	fmt.Println("This program is about age limit of movie .")
	fmt.Println()
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Print("Enter the day of the week: ")
	fmt.Scanln(&day)

	// detect
	if age < 5 || age > 95 {
		// free
		fmt.Println("You get in for free.")
	} else if (day == "tuesday" || day == "thursday") || (age > 12 && age < 21) {
		// student pricing
		fmt.Println("You are eligible for student pricing.")
	} else {
		// regular price
		fmt.Println("You must pay regular price.")
	}

	fmt.Println("\nDone.")
}
