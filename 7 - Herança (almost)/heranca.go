package main

import "fmt"

type person struct {
	name string
	last_name string
	height uint8
	age uint8
}

type student struct {
	person
	course string
	university string
}

func main() {
	fmt.Println("Inheritance")

	personOne := person{"Guilherme", "Schons", 173, 22}
	fmt.Println(personOne)

	studentOne := student{personOne, "Sistemas de Informação", "SETREM"}
	fmt.Println(studentOne)
	fmt.Println(studentOne.name)
	fmt.Println(studentOne.course)
}
