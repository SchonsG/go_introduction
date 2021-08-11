package main

import "fmt"

func main() {
	// Arithmetic (Some rules)

	// Error because they are differente types
	// var number1 int32 = 1000
	// var number2 int64 = 100000000
	// soma := number1 + number2
	// fmt.Println(soma)

	var number1 int32 = 1000
	var number2 int32 = 1000
	soma := number1 + number2
	fmt.Println(soma)


	// Relations Operators
	fmt.Println("----------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)


	// Logic Operators
	fmt.Println("----------------------")
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)


	// Unary Operators
	fmt.Println("----------------------")
	value := 10
	value++
	value += 10
	fmt.Println(value)

	value--
	value -= 10
	fmt.Println(value)

	value *= 3
	fmt.Println(value)
	value /= 3
	fmt.Println(value)
	value %= 3
	fmt.Println(value)


	// Ternary Operators -> don't exist.
	// text := nmumer > 5 ? "More than 5" : "Less than 5"
	// this is a wrong code. We don't have ternary operators

	// To do the same:

	var someText string
	someValue := 5

	if someValue > 5 {
		someText = "More than 5"
	} else {
		someText = "Less than 5, or equal"
	}

	fmt.Println(someText)
}