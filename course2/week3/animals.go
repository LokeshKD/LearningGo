// Program to respond to users input about an animal and and inquiry about its attributes.
// Attributes, in this case, are Food it eats, how to moves, sound it makes.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() string { return a.food }
func (a *Animal) Move() string { return a.locomotion }
func (a *Animal) Speak() string { return a.noise }

func main () {
	animals := map[string]Animal {
		"cow" : { "grass", "walk", "moo"},
		"bird": { "worms", "fly", "peep"},
		"snake": { "mice", "slither", "hsss"},
	}

	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nInput <animal name> and its <feature> on prompt \">\"")
		fmt.Println("Choices of Animals are [ cow, bird, snake]")
		fmt.Println("Choices of Feature are [ eat, move, speak]\n")

		fmt.Print("\n>")
		input.Scan()

		word := strings.Split(input.Text(), " ")
		animal, ok := animals[strings.ToLower(word[0])]
		if !ok {
			fmt.Println("Unidentified Animal <" + word[0] + ">\n")
			continue
		}

		switch strings.ToLower(word[1]) {
		case "eat":
			fmt.Printf("Animal %s Eats %s\n", word[0], animal.Eat())
		case "move":
			fmt.Printf("Animal %s Moves by %s\n", word[0], animal.Move())
		case "speak":
			fmt.Printf("Animal %s Sounds %s\n", word[0], animal.Speak())
		default:
			fmt.Println("Unidentified Feature <" + word[1] + "> Try again\n")
		}
	}
}
