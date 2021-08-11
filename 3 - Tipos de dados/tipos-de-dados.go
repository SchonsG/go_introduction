package main

import (
	"errors"
	"fmt"
)

func main() {

	// Basic types.

	// int (4 types).
	// int8, int16, int32, int64

	var number1 int64 = 100000000000
	fmt.Println(number1)

	// Example overflow.
	// var number2 int8 = 100000000000
	// fmt.Println(number2)

	// Without value in 'int', the bits usage from int is the same of your computer.
	var number3 int = 100000000000
	fmt.Println(number3)

	// unsigned int -> only 0 or positive numbers.
	var number4 uint32 = 1000
	fmt.Println(number4)

	// Alias
	// INT32 = RUNE (General is used for work with ASCII table characters).
	var number5 rune = 123456
	fmt.Println(number5)

	// BYTE = UINT8
	var number6 byte = 123
	fmt.Println(number6)

	// float
	// float32, float64

	var realNumber float32 = 123.45
	fmt.Println(realNumber)

	// Strings need be double quots
	// Char will be a number
	char := 'B'
	fmt.Println(char)
	// Exit: 66 (representation fo ASCII table), wanna be a INT, NOT A CHAR!


	// Error is a type.
	var problem1 error // default value is <nil>
	fmt.Println(problem1)

	var problem2 error = errors.New("Intern error")
	fmt.Println(problem2)

}