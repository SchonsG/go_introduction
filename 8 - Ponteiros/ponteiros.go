package main

import "fmt"

func main() {
	var variableOne int = 10
	var variableTwo int = variableOne // Copy, not reference
	fmt.Println(variableOne, variableTwo) // exit: 10, 10

	variableOne++
	fmt.Println(variableOne, variableTwo) // exit: 11, 10

	// Pointer
	var variableThree int = 10
	var pointer *int = &variableThree // & means variable reference
	fmt.Println(pointer) // Memory address
	fmt.Println(*pointer) // Dereferencing
	fmt.Println(variableOne, *pointer) // exit: 10, 10

	variableThree++
	fmt.Println(variableOne, *pointer) // exit: 11, 11
}