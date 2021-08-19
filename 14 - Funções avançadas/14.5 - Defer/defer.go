package main

import "fmt"


func functionOne() {
	fmt.Println("running function one")
}

func functionTwo() {
	fmt.Println("running function two")
}

func approveStudent(n1, n2 float32) bool {

	media := (n1 + n2) / 2

	defer fmt.Println("Avarage calculated. Result will be shown")

	if media > 6 {
		return true
	}

	return false
}

func main() {
	// Defer the function until the last moment of executation
	defer functionOne()

	functionTwo()

	fmt.Println(approveStudent(7,8))
}