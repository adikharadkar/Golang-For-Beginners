// Logical Operators

package main

import (
	"fmt"
)

func main() {
	num1 := true
	num2 := false
	num3 := true
	num4 := false
	
	result1 := num1 && num2
	fmt.Println("true && false =", result1)

	result2 := num1 && num3
	fmt.Println("true && true =", result2)

	result3 := num2 && num4
	fmt.Println("false && false =", result3)

	result4 := num1 || num2
	fmt.Println("true || false =", result4)

	result5 := num1 || num3
	fmt.Println("true || true =", result5)

	result6 := num2 || num4
	fmt.Println("false || false =", result6)

	result7 := !num1
	fmt.Println("!true =", result7)

	result8 := !num2
	fmt.Println("!false =", result8)
}
