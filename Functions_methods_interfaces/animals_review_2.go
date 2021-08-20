package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a Animal) getInfo(action string) {
	switch action {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func main() {
	cow := Animal{"grass",  "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Println("> Type your animal choice, type 'quit' to exit")
		fmt.Print("> ")

		input, _ := consoleReader.ReadString('\n')

		input = strings.ToLower(input)
		inputArr := strings.Fields(input)

		switch inputArr[0] {
		case "cow":
			cow.getInfo(inputArr[1])
		case "snake":
			snake.getInfo(inputArr[1])
		case "bird":
			bird.getInfo(inputArr[1])
		case "quit":
			fmt.Println("Good bye!")
			os.Exit(0)
		default:
			fmt.Println("No animal available")
		}
	}

}