package main

import (
	"fmt"

	"modulo/auxiliar"
)

func main() {
	fmt.Println("Writing from main function!")

	// Function imported from other package.
	auxiliar.Write()
}
