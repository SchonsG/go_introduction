package main

import (
	"fmt"

	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main function!")

	// Function imported from other package.
	auxiliar.Write()

	erro := checkmail.ValidateFormat("schons@gmail.com")
	fmt.Println(erro)
}
