package main

import (
	"reflect"
	"fmt"
)

func main() {
	fmt.Println("-------------------- Creating arrays")
	var arrayOne[5] string
	arrayOne[0] = "First position"
	fmt.Println(arrayOne)

	arrayTwo := [2] string {"First position", "Second position"}
	fmt.Println(arrayTwo)

	// Set array size with len of items sets (keep fix positions).
	arrayThree := [...] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arrayThree)


	// Slice (SLICE IS NOT AN ARRAY)
	fmt.Println("-------------------- Creating slice")
	sliceOne := []int {1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sliceOne)

	fmt.Println(reflect.TypeOf(sliceOne)) // exit: []int
	fmt.Println(reflect.TypeOf(arrayThree)) // exit: [10]int


	// append create a new Slice with the new item, and return its.
	fmt.Println("-------------------- Append slice")
	sliceTwo := append(sliceOne, 8)
	fmt.Println(sliceTwo)


	// Create a 'slice' from an array.
	fmt.Println("-------------------- Slice from array")
	sliceThree := arrayThree[:3]
	fmt.Println(sliceThree)

	// Slice from array, have the same behavior of pointer.
	fmt.Println("-------------------- Behavior slice")
	fmt.Println(arrayThree)
	fmt.Println(sliceThree)
	arrayThree[1] = 100
	fmt.Println(arrayThree)
	fmt.Println(sliceThree)
}