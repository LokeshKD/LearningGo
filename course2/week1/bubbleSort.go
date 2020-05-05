// Program to implment Bubble Sort in an array of 10 Integers.

package main

import (
	"fmt"
)

func main() {
		fmt.Printf("Input a sequence of 10 Integers, on prompt, one by one\n")

		var i, in int
		arr := make([]int, 0, 10)

		for i < 10 {
			fmt.Printf("Input element %d:", i)
			fmt.Scan(&in)
			arr = append(arr, in)
			i++
		}
		fmt.Println("Input array is...")
		fmt.Println(arr)
		bubbleSort(arr)
		fmt.Println("Sorted array is...")
		fmt.Println(arr)
}

func bubbleSort(arr []int) {
		length := len(arr)-1
		for i := 0; i < length; i++ {
			for j := 0; j < length - i; j++ {
				if arr[j] > arr[j+1] {
					swap(arr, j)
				}
				fmt.Println(arr)
			}
		}
}

func swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}


