package main

import (
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worm")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	// userPrompt here
	var input1 string
	var input2 string
	var input3 string
	var zoo = make(map[string]Animal)
	for {
		fmt.Println(strings.Repeat("*", 50))
		fmt.Println(strings.Repeat("*", 50))
		fmt.Println("Usage: newanimal - define a new animal. Types: cow, bird or snake")
		fmt.Println("Example: newanimal <name> <type>")
		fmt.Println("Usage: query - query a specific animal's action by the animal's name. Actions: move, eat or speak")
		fmt.Println("Example: query <name> <action>")
		fmt.Println("Type 'exit' to exit the application")
		fmt.Println(strings.Repeat("*", 50))
		fmt.Println(strings.Repeat("*", 50))
		fmt.Print("> ")
		fmt.Scanln(&input1, &input2, &input3)
		switch strings.ToLower(input1) {
		case "exit":
			os.Exit(1)
		case "newanimal":
			switch input3 {
			case "cow":
				zoo[input2] = Cow{}
				fmt.Println(strings.Repeat("*", 40))
				fmt.Printf("You added a %s named %s to your zoo\n", input3, input2)
				fmt.Println(strings.Repeat("*", 40))
				fmt.Println("So far your zoo contains: ", zoo)
			case "bird":
				zoo[input2] = Bird{}
				fmt.Println(strings.Repeat("*", 40))
				fmt.Printf("You added a %s named %s to your zoo\n", input3, input2)
				fmt.Println(strings.Repeat("*", 40))
				fmt.Println("So far your zoo contains: ", zoo)
			case "snake":
				zoo[input2] = Snake{}
				fmt.Println(strings.Repeat("*", 40))
				fmt.Printf("You added a %s named %s to your zoo\n", input3, input2)
				fmt.Println(strings.Repeat("*", 40))
				fmt.Println("So far your zoo contains: ", zoo)
			}
		case "query":
			switch input3 {
			case "eat":
				zoo[input2].Eat()
			case "move":
				zoo[input2].Move()
			case "speak":
				zoo[input2].Speak()
			}
		default:
			fmt.Println("Please enter the correct options")
			continue
		}
	}
}
