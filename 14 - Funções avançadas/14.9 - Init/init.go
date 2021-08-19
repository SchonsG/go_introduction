package main

import "fmt"

var test string

func main() {
	fmt.Println("Main Function")
	fmt.Println(test)
}

// Firt to be executed - Just one per archive
func init() {
	fmt.Println("Init Function")
	test = "Hello World"
}
