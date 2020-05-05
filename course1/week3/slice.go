//Program to take integers as input, store them and sort them in slices.

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main () {
	var in string

	slice := make([]int, 0, 3)

	for {
		fmt.Printf("Input an Integer to store in a Sorted Slice. X to exit :")

		_, err := fmt.Scan(&in)
		if err != nil {
			fmt.Println("Invalid Input, Try again...")
			continue
		}
		if in == "X" { return }

		intVal, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println("Invalid Input, Try again...")
			continue
		}
		slice = append(slice, intVal)

		sort.Ints(slice)
		fmt.Printf("Sorted Slice is %d \n", slice)
	}
}
