//Program to use go routines to sort a set of integers of an array partioned into 4 equal parts.
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Provide a sequence of Integers separated by a Space: ")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := strings.Split(input.Text(), " ")
	if len(line) < 4 {
		fmt.Println("Too few integers to sort. Exiting")
		return
	}

	ints := make([]int, 0)

	for _, item := range line {
		num, _ := strconv.Atoi(item)
		ints = append(ints, num)
	}

	fmt.Println("UnSorted Array is...")
	fmt.Println(ints)

	slice_len := len(ints)/4
	for i:=0; i<4; i++ {

		begin := i * slice_len
		end := (i+1) * slice_len
		if end > len(ints) {
			end = len(ints)
		}

		wg.Add(1)
		go func(ints []int) {
			sort.Ints(ints)
			fmt.Println(ints)
			wg.Done()
		}(ints[begin:end])
	}
	wg.Wait()
	fmt.Println("Sorted(Concurrently) Slices are...")
	fmt.Println(ints)

	sort.Ints(ints)
	fmt.Println("Sorted(Concurrently) Array is...")
	fmt.Println(ints)
}
