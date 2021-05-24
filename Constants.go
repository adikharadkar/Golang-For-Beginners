package main

import (
	"fmt"
)

func main() {

	// We use 'const' keyword to declare the constants
	// Initializing the constant on the same line by assigning a value '12'.
	const A int = 12
	fmt.Println("Value =", A)		// Printing the value
	fmt.Printf("Type = %T", A)		// Printing the type

	// Assigning a new value to the constant 'A'
	// But this will throw an error.
	// Because we cannot change the value of a constant
	A = 14
}
