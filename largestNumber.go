package main

import "fmt"

func main() {
	var (
		counter int
		largest int
		number  int
	)

	largest = 0
	number = 0

	for counter = 0; counter < 10; counter++ {
		fmt.Print("Enter Number: ")
		fmt.Scanln(&number)
		if number > largest {
			largest = number
		}
	}
	fmt.Print("Largest number: ", largest)
}
