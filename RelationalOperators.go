// Relational operators are used to compare two variables

package main

import (
	"fmt"
)

func main() {
	// Declaring two integer variables
	num1 := 20
	num2 := 10

	// Checking whether both the numbers are equal or not
	// We use EqualTo (==) operator
	// Declaring the integer varable to store the result
	// If both the numbers are equal, then the result will be 'true', otherwise it will be 'false'
	result1 := num1 == num2
	fmt.Println("Result 1 =", result1) // Printing the result

	// Checking if both the numbers are not equal to each other
	// We use NotEqualTo (!=) operator
	// Declaring the integer varable to store the result
	// If both the numbers are not equal, then the result will be 'true', otherwise it will be 'false'
	result2 := num1 != num2
	fmt.Println("Result 2 =", result2) // Printing the result

	// Checking whether the first number is greater than the second number or not
	// We use GreaterThan (>) operator
	// Declaring the integer varable to store the result
	// If first number is grater than second number, then the result will be 'true', otherwise it will be 'false'
	result3 := num1 > num2
	fmt.Println("Result 3 =", result3) // Printing the result

	// Checking whether the first number is less than the second number or not
	// We use LessThan (>) operator
	// Declaring the integer varable to store the result
	// If first number is less than second number, then the result will be 'true', otherwise it will be 'false'
	result4 := num1 < num2
	fmt.Println("Result 4 =", result4) // Printing the result

	// Checking whether the first number is greater than or equal to the second number or not
	// We use GreaterThanOrEqual (>=) operator
	// Declaring the integer varable to store the result
	// If first number is grater than or equal to the second number, then the result will be 'true', otherwise it will be 'false'
	result5 := num1 >= num2
	fmt.Println("Result 5 =", result5) // Printing the result

	// Checking whether the first number is less than or equal to the second number or not
	// We use LessThanOrEqual (<=) operator
	// Declaring the integer varable to store the result
	// If first number is less than or equal to the second number, then the result will be 'true', otherwise it will be 'false'
	result6 := num1 <= num2
	fmt.Println("Result 6 =", result6) // Printing the result
}
