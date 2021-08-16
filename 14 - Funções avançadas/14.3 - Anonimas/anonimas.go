package main

import "fmt"

func main() {
	message := func(text string) string {
		return "Hello " + text
	}("World!")

	fmt.Println(message)
}