package main

import "fmt"

func main() {
	var (
		numberOne int
		numberTwo int
	)
	fmt.Print("Enter first number: ")
	fmt.Scanln(&numberOne)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&numberTwo)

	if numberOne == numberTwo {
		fmt.Println(0)
	} else if numberOne > numberTwo {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}
