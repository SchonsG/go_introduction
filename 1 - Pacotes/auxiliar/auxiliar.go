package auxiliar

import "fmt"

// Export a function called Write that write a message in console.
func Write() {
	fmt.Println("Writing from Auxiliar package")

	// Function from 'auxiliar2' archive, this archive is from the same package.
	escrever2()
}