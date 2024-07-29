package main

import "fmt"

func main() {
	var (
		counter       int
		largest       int
		secondLargest int
		number        int
	)

	largest = 0
	secondLargest = 0

	for counter = 0; counter < 10; counter++ {

		fmt.Print("Enter number: ")

		fmt.Scanln(&number)

		if number > largest {
			secondLargest = largest
			largest = number
		} else if number > secondLargest {
			secondLargest = number
		}

	}

	fmt.Println("Largest number: ", largest)
	fmt.Println("Second largest number: ", secondLargest)
}
