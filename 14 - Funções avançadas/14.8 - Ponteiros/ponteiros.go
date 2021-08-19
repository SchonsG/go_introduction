package main

import "fmt"


func reverseSignWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	fmt.Println(number)
	reverseSignWithPointer(&number)
	fmt.Println(number)
}