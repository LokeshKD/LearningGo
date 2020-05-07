// Program to demonstrate Race condition with go routines.

package main

import (
	"fmt"
	"time"
)

// Function that incements the value at an index i by 1 from the next index value.
func fn(arr []int, i int) {
	arr[i] = arr[i+1] + 1
	fmt.Println(arr[:7])
}

func main() {
	arr := []int {1,2,3,4,5,6,7,8}

	// Below we have 7 go routines with concurrent execution and access to
	// shared slice, without any mutual exclusion, we have a data RACE condition.
	// The array will be printed 7 times and with every run the contents will be different.
	// Values will be different even within the array, on every output.
	// This is because execution of all the 7 go routines is not deterministic.
	for i:=0; i < cap(arr)-1; i++ {
		go fn(arr, i)
	}

	time.Sleep(1 * time.Millisecond)// Wait a little time.
}
