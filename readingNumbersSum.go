package main

import "fmt"

func main() {
	var (
		numberChosen int
		sum          int
		number       int
	)
	fmt.Print("Enter a chosen number: ")
	fmt.Scanln(&numberChosen)

	for sum < numberChosen {

		fmt.Print("Enter a number: ")
		fmt.Scanln(&number)
		sum += number
	}

	fmt.Printf("Sum of numbers is greater than or equal to the chosen number: %d\n", sum)
}
