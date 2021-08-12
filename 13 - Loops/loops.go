package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	// fmt.Println(j) -> doesn't work, because j is a scope variable


	// Clasule 'range' is for when you need iterate something
	names := [3]string {"Guilherme", "Tiago", "Rafael"}

	// index and name
	for index, name := range names {
		fmt.Println(index, name)
	}

	// just name
	for _, name := range names {
		fmt.Println(name)
	}

	for index, letter := range "WORD" {
		fmt.Println(index, letter) // letter = ASCII table
		fmt.Println(index, string(letter))
	}

	user := map[string] string {
		"name": "Guilherme",
		"last_name": "Schons",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	// for in structs doesn't works!


	// For withtou conditions, is a infinite loop
	// for {

	// }

}