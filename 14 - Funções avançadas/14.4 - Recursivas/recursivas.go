package main

import "fmt"


func fibonacci(pos uint) uint {
	if pos <= 1 {
		return pos
	}

	return fibonacci(pos - 2) + fibonacci(pos - 1)
}

func main() {
	position := uint(10)
	fmt.Println(fibonacci(position))
}