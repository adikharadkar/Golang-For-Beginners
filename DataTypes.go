package main

import  (
	"fmt"
)

func main () {

	// Declaring a variable of 'string' data type.
	var channel string = "Brainstorm Codings"

	// Printing the value stored in the variable i.e. Brainstorm Codings.
	fmt.Println("Value stored in channel variable is", channel)

	// Printing the type of the variable i.e. string.
	fmt.Printf("Type of channel variable is %T", channel)

	// Declaring a variable of 'integer' data type.
	var number int = 90

	// Printing the value stored in the variable i.e. number
	fmt.Println("Value stored in number variable is", number)

	// Printing the type of the variable i.e. int.
	fmt.Printf("Type of number variable is %T", number)

	// Declaring a variable of 'float64' type.
	var number2 float64 = 12.5

	// Printing the value stored in the variable i.e. 12.5
	fmt.Println("Value stored in number2 variable is", number2)

	// Printing the type of the variable i.e. float64.
	fmt.Printf("Type of number2 variable is %T", number2)
}
