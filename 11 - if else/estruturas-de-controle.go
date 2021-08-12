package main

import "fmt"

func main() {

	number := 10

	if number > 15 {
		fmt.Println("More than 15")
	} else {
		fmt.Println("Less or equal than 15")
	}

	// variable init (otherNumber)... just runs in the scope.
	if otherNumber := 1; otherNumber > 0 {
		fmt.Println("More than zero")
	}
}