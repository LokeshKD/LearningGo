// Program to read from a file, first and last names of a person(s) and print them o Stdout.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname, lname  string
}

func main() {
	names := make([]name, 0, 4)

	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("Name of File to read data from:")
	input.Scan()
	filename := input.Text()

	// Read the file contents.
	file, err := os.Open(filename)
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 3 { continue } // Shortest name could be "a b"
		words := strings.Split(line, " ")
		p := name{words[0], words[1]}
		names = append(names, p)
	}

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}

	fmt.Println("\nNames of the People are: {First Last}")
	for _, p := range names {
		fmt.Println(p)
	}
}

