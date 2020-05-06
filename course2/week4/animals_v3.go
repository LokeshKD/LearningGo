// Program to respond to users input about an animal and and inquiry about its properties.
// Properties, in this case, are Food it eats, how it moves, sound it makes.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {name, food, locomotion, noise string}
type Bird struct {name, food, locomotion, noise string}
type Snake struct {name, food, locomotion, noise string}

func MakeCow(name string) *Cow { return &Cow{name, "grass", "walk", "moo"} }
func MakeBird(name string) *Bird { return &Bird{name, "worms", "fly", "peep"} }
func MakeSnake(name string) *Snake { return &Snake{name, "mice", "slither", "hsss"} }

func (a *Cow) Eat() { fmt.Printf("%s Eats %s\n",a.name, a.food) }
func (a *Cow) Move() { fmt.Printf("%s moves by %s\n",a.name, a.locomotion) }
func (a *Cow) Speak() { fmt.Printf("%s says %s\n",a.name, a.noise) }

func (a *Bird) Eat() { fmt.Printf("%s Eats %s\n",a.name, a.food) }
func (a *Bird) Move() { fmt.Printf("%s moves by %s\n",a.name, a.locomotion) }
func (a *Bird) Speak() { fmt.Printf("%s says %s\n",a.name, a.noise) }

func (a *Snake) Eat() { fmt.Printf("%s Eats %s\n",a.name, a.food) }
func (a *Snake) Move() { fmt.Printf("%s moves by %s\n",a.name, a.locomotion) }
func (a *Snake) Speak() { fmt.Printf("%s says %s\n",a.name, a.noise) }


func main () {

	animals := make(map[string]Animal)

	fmt.Println("\nTwo commands that I work with [newanimal query] exit/quit to terminate")
	fmt.Println("newanimal <name of animal> Type[cow bird snake]")
	fmt.Println("query <name of animal> Property[Eat Move Speak]")

	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n>")
		input.Scan()

		word := strings.Split(input.Text(), " ")
		cmd := strings.ToLower(word[0])

		if cmd == "quit" || cmd == "exit" { break }

		if len(word) < 3 {
			fmt.Println("Too few arguments, Try again...")
			continue
		}

		switch cmd {
		case "newanimal":
			NewAnimal(animals, word[1:3])
		case "query":
			NewQuery(animals, word[1:3])
		default:
			fmt.Println("Unknown command <" + word[0] + "> Try again...\n")
		}
	}
}

//Create a new animal based on the type and add it to a map of that type.
// We also store it in types map for easy retrieval later in Query phase.
func NewAnimal(animals map[string]Animal, words []string) {

	name := strings.ToLower(words[0])

	switch strings.ToLower(words[1]) {
	case "cow":
		animals[name] = MakeCow(name)
	case "bird":
		animals[name] = MakeBird(name)
	case "snake":
		animals[name] = MakeSnake(name)
	default:
		fmt.Printf("Unknown Animal Type <%s> Cannot Create it\n", words[1])
		return
	}
	fmt.Println("Created it!")
}

// Query the database for a specific feature of an animal created.
func NewQuery(animals map[string]Animal, words []string) {

	var animal Animal

	name := strings.ToLower(words[0])
	act := strings.ToLower(words[1])

	animal, ok := animals[name]
	if !ok {
		fmt.Printf("Could not find an animal with Name %s\n", words[0])
		return
	}

	ProcessQuery(animal, act)
}

// Call the specific Feature with Animal Interface.
func ProcessQuery(animal Animal, act string) {

	switch act {
		case "eat": animal.Eat()
		case "move": animal.Move()
		case "speak": animal.Speak()
		default:
			fmt.Printf("Unknown Feature requested %s\n", act)
	}
}

