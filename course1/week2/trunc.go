// A program to print an integer with a user input floating point number.
package main

import "fmt"

func main() {
	var in float64

	fmt.Printf("Provide a Floating Point Number to be converted to Integer:")
	_, err := fmt.Scan(&in)
	if err!= nil {
		fmt.Printf("\nInput Provided is not a Floating Point Number\n")
		return
	}
	fmt.Printf("\nFloating Point Number is %f and Its Integer Value is %d \n", in, int64(in))
}
