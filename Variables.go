package main

import (
	"fmt"
)

func main() {

	// Variable Declaration using 'var keyword'
	var a int

	// Initializing variable by assigning a value
	a = 12
	fmt.Println("Value =", a)		// Printing value of the variable
	fmt.Printf("Type = %T", a)		// Printing type of the variable

	// Declaring the variable using 'var keyword' and initializing on the same line
	var b string = "Brainstorm Codings"
	fmt.Println("Value =", b)		// Printing value of the variable
	fmt.Printf("Type = %T", b)		// Printing type of the variable

	// This is the 'Short-Hand Declaration' method to declare the variable
	// We use ':=' short-hand declaration operator.
	// Here we do not need to define the data type of the variable.
	c := 12.45
	fmt.Println("Value =", c)		// Printing value of the variable
	fmt.Printf("Type = %T", c)		// Printing type of the variable
}
