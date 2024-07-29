package main

import "fmt"

func main() {
	var (
		accountNumber    int
		beginningBalance int
		totalCharges     int
		totalCredits     int
		creditLimit      int
	)

	for accountNumber != 1 {
		fmt.Println("Enter account number")
		fmt.Scanln(&accountNumber)
		fmt.Println("Enter beginning balance")
		fmt.Scanln(&beginningBalance)
		fmt.Println("Enter total charges:")
		fmt.Scanln(&totalCharges)
		fmt.Println("Enter total credits:")
		fmt.Scanln(&totalCredits)
		fmt.Println("Enter credit limits")
		fmt.Scanln(&creditLimit)

		newBalance := beginningBalance + (totalCharges - totalCredits)

		fmt.Println("New Balance: ", newBalance)
		if newBalance > creditLimit {
			fmt.Println("Credit Limit exceeded")
		}

	}

}
