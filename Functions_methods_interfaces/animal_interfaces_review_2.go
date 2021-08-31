// Author: Konstantin Kutsner
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

// Cow

type Cow struct {
}

func (cow *Cow) Eat() string {
	return "grass"
}

func (cow *Cow) Move() string {
	return "walk"
}

func (cow *Cow) Speak() string {
	return "moo"
}

// Bird

type Bird struct {
}

func (bird *Bird) Eat() string {
	return "worms"
}

func (bird *Bird) Move() string {
	return "fly"
}

func (bird *Bird) Speak() string {
	return "peep"
}

// Snake

type Snake struct {
}

func (snake *Snake) Eat() string {
	return "mice"
}

func (snake *Snake) Move() string {
	return "slither"
}

func (snake *Snake) Speak() string {
	return "hsss"
}

func main() {

	// Let's define our nano database
	animals := map[string]Animal{}

	// Let's start our main loop
	const prompt = "\n> "
	fmt.Print(prompt)
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {

		// Get the line of text and check if need to exit on empty string
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		// Split the line into words
		words := strings.Split(line, " ")
		if len(words) != 3 {
			fmt.Print("Invalid format! Please try again!", prompt)
			continue
		}

		// Execute the requested command and pass the arguments
		switch words[0] {
		case "newanimal":
			NewAnimal(&animals, words[1], words[2])
		case "query":
			Query(&animals, words[1], words[2])
		default:
			fmt.Print("Unknown command! Please try again: newanimal or query!", prompt)
		}

		// Let's get another prompt
		fmt.Print(prompt)
	}

}

func NewAnimal(animals *map[string]Animal, animalName string, animalType string) {

	// Basics error handling
	if len(animalName) == 0 {
		fmt.Println("Animal name cannot be empty")
		return
	}

	animalType = strings.ToLower(animalType)
	switch animalType {
	case "cow":
		(*animals)[animalName] = &Cow{}
	case "bird":
		(*animals)[animalName] = &Bird{}
	case "snake":
		(*animals)[animalName] = &Snake{}
	default:
		fmt.Println("Unknown animal type")
		return
	}
	fmt.Println("It created!")
}

func Query(animals *map[string]Animal, animalName string, animalMethod string) {
	// Check unput data
	if len(animalName) == 0 {
		fmt.Println("Animal name cannot be empty")
		return
	}

	if len(animalMethod) == 0 {
		fmt.Println("Animal method cannot be empty")
		return
	}

	// Find the animal object by the name
	animal, found := (*animals)[animalName]
	if !found {
		fmt.Printf("Animal '%v' is not found!", animalName)
		return
	}

	// Call the requested method
	switch animalMethod {
	case "speak":
		fmt.Println(animal.Speak())
	case "eat":
		fmt.Println(animal.Eat())
	case "move":
		fmt.Println(animal.Move())
	default:
		fmt.Println("Method '%v' is not supported!")
	}
}
