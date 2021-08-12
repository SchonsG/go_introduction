package main

import (
	"fmt"
	"strconv"
)

// Variatic parameter need be unique in placed in the final assignature
func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func concatStringWithNumber(text string, numbers ...int) string {
	result := text
	for _, number := range numbers {
		result += strconv.Itoa(number)
	}
	return result
}

func main() {
	result := sum(1, 2, 3)
	fmt.Println(result)

	resultTwo := concatStringWithNumber("Hello ", 1, 2, 3)
	fmt.Println(resultTwo)
}