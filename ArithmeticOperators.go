// Arithmetic Operators

package main

import (
	"fmt"
)

func main() {
	// Declaring two integer variables
	num1 := 12
	num2 := 24

	// Performing addition on both numbers
	// Addition operator (+) is used to perform addition.
	// Declaring a variable to store addition of two variables
	addition := num1 + num2
	fmt.Println("Addition =", addition) // Printing the addition of two numbers

	// Performing multiplication on both numbers
	// Multiplication operator (*) is used to perform multiplication
	// Declaring an integer variable to store multiplication of two numbers
	multiplication := num1 * num2
	fmt.Println("Multiplication =", multiplication) // Printing te multiplication

	// Performing Subtraction on both numbers
	// Subtraction operator (-) is used to perform subtraction
	// Declaring an integer variable to store subtraction of two numbers
	subtraction := num2 - num1
	fmt.Println("Subtraction =", subtraction) // Printing the subtraction

	// Performing Division on both numbers
	// Division operator (/) is used to perform subtraction
	// Declaring an integer variable to store Division of two numbers
	devision := num2 / num1
	fmt.Println("Devision =", devision) // Printing the devision

	// Getting the remainder from the devision of two numbers
	// Modulus Operator (%) is used to get remainder
	// Declaring an integer variable to store the remainder
	mod := num2 % num1
	fmt.Println("Remainder =", mod) // Printing the remainder
}
