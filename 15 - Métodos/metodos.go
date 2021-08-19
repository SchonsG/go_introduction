package main

import ("fmt")


type user struct {
	name string
	age uint8
}

func (u user) save() {
	fmt.Printf("saving user %s...\n", u.name)
}

func (u *user) birth() {
	u.age++
}

func main() {
	userOne := user{"Guilherme", 22}
	fmt.Println(userOne)

	userOne.save()

	userOne.birth()
	fmt.Println(userOne)
}