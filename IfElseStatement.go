// if-else statement

package main

import (
	"fmt"
)

func main() {

	// Declaring integer variables
	num1 := 20
	num2 := 10

	// Creating an if-else condition.
	// This condtion will check whether 20 is equal to 10 or not.
	// If the condition is true, then the if block will get executed.
	// If the condition is false, then the else block will get executed.
	if num1 == num2 {
		fmt.Println("20 is equal to 10")
	} else {
		fmt.Println("20 is not equal to 10")
	}
}
