package main

import "fmt"

func main() {
	var (
		payPerWeek     float64 = 200
		commissionRate float64 = 0.09
		totalSales     float64 = 0
		numberOfItems  int
		itemPrice      float64
	)

	fmt.Print("Enter the number of items sold: ")
	fmt.Scanln(&numberOfItems)

	for count := 0; count < numberOfItems; count++ {

		fmt.Printf("Enter value of item %d: ", count+1)
		fmt.Scanln(&itemPrice)
		totalSales += itemPrice
	}

	var earnings float64 = payPerWeek + (totalSales * commissionRate)
	fmt.Printf("Total sales for the week: $%.2f\n", totalSales)
	fmt.Printf("Earnings for the week: $%.2f\n", earnings)
}
