// Program to convert a map[string]string into JSON format.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	person := make(map[string]string)
	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("Name of the Person:")
	input.Scan()
	person["name"] = input.Text()

	fmt.Printf("Address of the Person:")
	input.Scan()
	person["address"] = input.Text()

	data, err := json.MarshalIndent(person, "", " ")
	if err != nil {
		log.Fatalf("JSON Marshalling failed")
	}
	fmt.Printf("\nMarshaled JSON output is as below: \n%s\n", data)
}
