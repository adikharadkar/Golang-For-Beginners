/*
	Identifiers are used for identification purpose.
	Identifier is a user-given name to a variable, function, channel, array, or slices.
*/

/*
	Keyword are the special words in golang.
	They are reserved words and cannot be used as identifiers.
	In this program, we have following keywords: package, import, func, and var.
*/

// 'package' is a keyword.
package main

// 'import' is a keyword.
import (
	"fmt"
)

// We have 'func' which is a keyword.
func main () {
	// Here 'var' is a keyword and 'channel' is an identifier.
	var channel string = "Brainstorm Codings"
	fmt.Println(channel)
}
