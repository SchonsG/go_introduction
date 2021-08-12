package main

import "fmt"

type user struct {
	name string
	age uint8
	localization address
}

type address struct {
	state string
	city string
	cep uint32
}

func main() {
	// First form.
	var userOne user
	userOne.name = "Guilherme"
	userOne.age = 22

	var addressUserOne address
	addressUserOne.state = "Rio Grande do Sul"
	addressUserOne.city = "Três de Maio"
	addressUserOne.cep = 98910000

	// Attribution address to user
	userOne.localization = addressUserOne

	fmt.Println(userOne)

	// Second form.
	addressUserTwo := address{"Rio Grande do Sul", "Três de Maio", 9891000}
	userTwo := user{"Guilherme", 22, addressUserTwo}

	fmt.Println(userTwo)

	// Third form.
	addressUserThree := address{state: "RS"}
	userThree := user{name:"Guilherme", localization:addressUserThree}

	fmt.Println(userThree)
}