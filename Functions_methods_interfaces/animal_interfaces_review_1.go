package main

import (
	"fmt"
	"strings"
)

// newanimal Nanda cow
// newanimal Curly snake

// query Nanda eat
// query Nanda speak
// query Curly move

var p = fmt.Print

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

type Cow struct {
	name string
}

func (animal Bird) Eat() {
	p("worms")
}

func (animal Bird) Move() {
	p("fly")
}

func (animal Bird) Speak() {
	p("peep")
}

func (animal Cow) Eat() {
	p("grass")
}

func (animal Cow) Move() {
	p("walk")
}

func (animal Cow) Speak() {
	p("moo")
}

func (animal Snake) Eat() {
	p("mice")
}

func (animal Snake) Move() {
	p("slither")
}

func (animal Snake) Speak() {
	p("hsss")
}

func main() {

	animalMap := make(map[string]Animal, 3)

	for {
		command, arg1, arg2 := getInput()

		switch command {
		case "newanimal":
			newAnimal(arg1, arg2, animalMap)

		case "query":
			query(arg1, arg2, animalMap)

		default:
			fmt.Printf("Invalid command %q. Valid commands are 'newanimal' and 'query'.", command)

		}
	}
}

// query prints the info about an animal
func query(name, attribute string, animalMap map[string]Animal) {

	animal, found := animalMap[strings.ToUpper(name)]

	if found {

		switch attribute {
		case "eat":
			animal.Eat()

		case "move":
			animal.Move()

		case "speak":
			animal.Speak()

		default:
			fmt.Printf("Unknown attribute %q.", attribute)
		}

	} else {
		p(name, " not found.")
	}
}

// newAnimal creates an animal and saves the animal in a map so that it can be queried in future.
// The key of the map is string which is a concatenation of name of animal and type of animal.
func newAnimal(name, animalTypeInput string, animalMap map[string]Animal) {

	animalType := strings.ToUpper(animalTypeInput)
	var newAnimal Animal

	switch animalType {
	case "COW":
		newAnimal = Cow{name}

	case "SNAKE":
		newAnimal = Snake{name}

	case "BIRD":
		newAnimal = Bird{name}

	default:
		fmt.Printf("Unkonwn animal type %q.", animalTypeInput)
		return
	}

	_, found := animalMap[strings.ToUpper(name)]

	if found {
		fmt.Printf("Entry for %q which is a %q is already saved. Previous entry will be overwritten. ", name, animalTypeInput)
	}

	animalMap[strings.ToUpper(name)] = newAnimal

	p("Created it!")
}

func getInput() (string, string, string) {
	p("\n> ")
	var command, arg1, arg2 string
	fmt.Scan(&command, &arg1, &arg2)
	return command, arg1, arg2
}
