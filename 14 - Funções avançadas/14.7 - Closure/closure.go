package main

import "fmt"


func closure() func() {
	text := "Inside closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Inside main funtion"
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}