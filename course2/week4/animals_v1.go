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

type Cow struct{ name, food, locomotion, noise string }
type Bird struct{ name, food, locomotion, noise string }
type Snake struct{ name, food, locomotion, noise string }

func MakeCow(name string) *Cow {
	cow := Cow{name, "grass", "walk", "moo"}
	return &cow
}
func MakeBird(name string) *Bird {
	bird := Bird{name, "worms", "fly", "peep"}
	return &bird
}
func MakeSnake(name string) *Snake {
	snake := Snake{name, "mice", "slither", "hsss"}
	return &snake
}

func (a *Cow) Eat()   { fmt.Printf("%s Eats %s\n", a.name, a.food) }
func (a *Cow) Move()  { fmt.Printf("%s moves by %s\n", a.name, a.locomotion) }
func (a *Cow) Speak() { fmt.Printf("%s says %s\n", a.name, a.noise) }

func (a *Bird) Eat()   { fmt.Printf("%s Eats %s\n", a.name, a.food) }
func (a *Bird) Move()  { fmt.Printf("%s moves by %s\n", a.name, a.locomotion) }
func (a *Bird) Speak() { fmt.Printf("%s says %s\n", a.name, a.noise) }

func (a *Snake) Eat()   { fmt.Printf("%s Eats %s\n", a.name, a.food) }
func (a *Snake) Move()  { fmt.Printf("%s moves by %s\n", a.name, a.locomotion) }
func (a *Snake) Speak() { fmt.Printf("%s says %s\n", a.name, a.noise) }

// Maps to store all the typed animals with their name as key.
var cows map[string]*Cow
var birds map[string]*Bird
var snakes map[string]*Snake

func main() {

	cows = make(map[string]*Cow)
	birds = make(map[string]*Bird)
	snakes = make(map[string]*Snake)

	fmt.Println("\n Two commands that I work with")
	fmt.Println("newanimal <name of animal> Type[cow bird snake]")
	fmt.Println("query <name of animal> Property[Eat Move Speak]")

	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n>")
		input.Scan()

		word := strings.Split(input.Text(), " ")

		switch strings.ToLower(word[0]) {
		case "newanimal":
			NewAnimal(word[1:3])
		case "query":
			NewQuery(word[1:3])
		default:
			fmt.Println("Unknown command <" + word[0] + "> Try again...\n")
		}
	}
}

//Create a new animal based on the type and add it to a map of that type.
func NewAnimal(words []string) {

	name := strings.ToLower(words[0])

	switch strings.ToLower(words[1]) {
	case "cow":
		cows[name] = MakeCow(name)
	case "bird":
		birds[name] = MakeBird(name)
	case "snake":
		snakes[name] = MakeSnake(name)
	default:
		fmt.Printf("Unknown Animal Type <%s> Cannot Create it\n", words[1])
		return
	}
	fmt.Println("Created it!")
}

func NewQuery(words []string) {

	var animal Animal
	var ok bool
	name := strings.ToLower(words[0])
	act := strings.ToLower(words[1])

	animal, ok = cows[name]
	if ok {
		ProcessQuery(animal, act)
		return
	}
	animal, ok = birds[name]
	if ok {
		ProcessQuery(animal, act)
		return
	}
	animal, ok = snakes[name]
	if ok {
		ProcessQuery(animal, act)
		return
	}
	fmt.Printf("Could not find an animal with Name %s\n", words[0])
}

func ProcessQuery(animal Animal, act string) {

	switch act {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Printf("Unknown Feature requested %s\n", act)
	}
}
