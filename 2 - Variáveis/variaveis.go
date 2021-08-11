// When you declare a variable, you need use or it's gonna crash!

package main

import "fmt"

func main() {
	// Explicit var --------------------------------------

	// First form.
	var name string = "Guilherme"
	fmt.Println(name)

	// Second form.
	var (
		estate string = "Rio Grande do Sul"
		city string = "Três de Maio"
	)
	fmt.Println(estate)
	fmt.Println(city)

	// Implicit var --------------------------------------

	// First form.
	occupation := "Developer"
	fmt.Println(occupation)

	// Second form.
	skillOne, skillTwo := "Go", "Python"
	fmt.Println(skillOne)
	fmt.Println(skillTwo)

	// ---------------------------------------------------

	// Great tip, change the vars values :)
	skillOne, skillTwo = skillTwo, skillOne
	fmt.Println(skillOne)
	fmt.Println(skillTwo)
}