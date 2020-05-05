// Program to search for certain charactes in an input string.
// Prints Found, if found and Not Found, otherwise.

package main

import "fmt"
import "strings"

func main() {
	var in string

	fmt.Printf("Input a string for valuation:")
	fmt.Scan(&in)
	in = strings.ToLower(in)
	if strings.Contains(in, "a") && strings.HasPrefix(in, "i") && strings.HasSuffix(in, "n") {
		fmt.Printf("\nFound!\n")
		return
	}
	fmt.Printf("\nNot Found!\n")
}
