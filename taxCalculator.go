package main

import "fmt"

func main() {
	var (
		numberOfCitizens int
		nameOfCitizen    string
		earnings         float64
		tax              float64
	)
	fmt.Print("Enter number of citizens: ")
	fmt.Scanln(&numberOfCitizens)

	for count := 0; count < numberOfCitizens; count++ {

		fmt.Println("Enter citizen's name: ")
		fmt.Scanln(&nameOfCitizen)
		fmt.Printf("Enter %s's earnings: ", nameOfCitizen)
		fmt.Scanln(&earnings)

		if earnings <= 30000 {
			tax = earnings * 0.15
		} else {
			tax = earnings * 0.20
		}

		fmt.Printf("%s's total tax is: $%.2f\n", nameOfCitizen, tax)
	}
}
