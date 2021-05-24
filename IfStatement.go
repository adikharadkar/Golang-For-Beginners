// if statement

package main

import (
	"fmt"
)

func main() {
  
  // Declaring variables
	num1 := 20
	num2 := 10

  // Creating an if statement to check whether 20 is equal to 10 or not.
  // If the condition is true, then the if block will get executed.
	if num1 == num2 {
		fmt.Println("20 is equal to 10")
	}

	if num1 != num2 {
		fmt.Println("20 is not equal to 10")
	}

	if num1 >= num2 {
		fmt.Println("20 is greater than or equal to 10")
	}

	if num1 <= num2 {
		fmt.Println("20 is less than or equal to 10")
	}
}
