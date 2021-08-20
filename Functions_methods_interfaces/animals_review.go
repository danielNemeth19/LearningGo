package main

/*
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the three animals and their associated data which should be hard-coded into your program.
Animal	Food eaten	Locomotion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.


*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	myConsoleReader := bufio.NewReader(os.Stdin)

	var development bool = false

	animalMap := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	fmt.Printf("\nWelcome to animals.\n\n")

	for {
		fmt.Printf("\nEnter the animal name (cow, bird, or snake) and query (eat, move, or speak) separated by spaces (or X to stop) > ")
		lineString, err := myConsoleReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if (strings.Contains(lineString, "X")) ||
			(strings.Contains(lineString, "x")) {
			break
		}
		lineString = strings.Replace(lineString, "\r", "", -1)
		lineString = strings.Replace(lineString, "\n", "", -1)
		if development {
			fmt.Printf("\nNew line has content %s:\n", lineString)
		}
		lineParts := strings.Split(lineString, " ")

		if len(lineParts) != 2 {
			fmt.Println("Remember to enter just two words: the animal name and your query.")
			continue
		}

		animal := lineParts[0]
		query := lineParts[1]

		if development {
			fmt.Printf("\nYour animal is %s, and you are asking about %s.", animal, query)
		}

		value, ok := animalMap[animal]
		if ok && development {
			fmt.Printf("\nWhat we know about the animal %s (eat, move, or speak): %s\n", animal, value)
		}
		if !ok {
			fmt.Printf("\nSorry, we do not have information about an animal called %s ", animal)
		}

		switch query {
		case "eat":
			value.Eat()
		case "move":
			value.Move()
		case "speak":
			value.Speak()
		default:
			fmt.Printf("Sorry. I don't know about %s ", query)
		}

	}

	fmt.Printf("\nHave a good day.\n\n")
}
