package main

import "fmt"

func main() {
	var (
		miles          int
		gallons        int
		totalMiles     int
		totalGallons   int
		milesPerGallon float64
	)

	totalMiles = 0
	totalGallons = 0
	milesPerGallon = 0

	for miles != -1 {
		fmt.Println("Enter miles driven(press -1 to quit): ")
		fmt.Scanln(&miles)

		fmt.Println("Enter gallons used: ")
		fmt.Scanln(&gallons)

		totalMiles += miles
		totalGallons += gallons

		milesPerGallon = float64(totalMiles) / float64(totalGallons)

		fmt.Println(milesPerGallon)

	}

}
