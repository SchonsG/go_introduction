package main

import (
	"errors"
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// More than 1 return
func mathCalcs(n1, n2 int8) (int8, int8) {
	return n1 + n2, n1 - n2
}

func compareStringsLenght(stringOne, stringTwo string) (error, bool) {
	if stringOne == "" || stringTwo == "" {
		var error = errors.New("Empty value in parameters")
		return error, false
	}

	if len(stringOne) == len(stringTwo) {
		return nil, true
	}

	return nil, false
}

func main() {
	result := somar(1,1)
	fmt.Println(result)

	var function = func(txt string) string {
		fmt.Println(txt)
		return "It Works!"
	}

	result_function := function("Hello")
	fmt.Println(result_function)

	result_add, result_sub := mathCalcs(1,1)
	fmt.Println(result_add)
	fmt.Println(result_sub)

	_, isEqual := compareStringsLenght("abc", "")
	fmt.Println(isEqual)
}
