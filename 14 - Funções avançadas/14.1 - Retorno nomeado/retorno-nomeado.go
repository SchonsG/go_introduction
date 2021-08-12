package main

import (
	"fmt"
)

// Nomed return
func mathCalcs(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	resultOne, resultTwo := mathCalcs(10, 20)
	fmt.Println(resultOne, resultTwo)
}