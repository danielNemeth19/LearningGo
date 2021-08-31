package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	// I know in fact this field is not needed, but it would be nice to store the name inside an animal
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

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
	animals := make(map[string]Animal)

	var command, name, arg string

	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&command, &name, &arg)
		if err == nil {
			switch command {
			case "newanimal":
				newAnimal(name, arg, &animals)
			case "query":
				queryAnimal(name, arg, &animals)
			}
		}
	}
}

func newAnimal(name string, species string, animals *map[string]Animal) {
	switch species {
	case "cow":
		(*animals)[name] = Cow{name}
	case "bird":
		(*animals)[name] = Bird{name}
	case "snake":
		(*animals)[name] = Snake{name}
	default:
		return
	}
	fmt.Println("Created it!")
}

func queryAnimal(name string, action string, animals *map[string]Animal) {
	switch action {
	case "eat":
		(*animals)[name].Eat()
	case "move":
		(*animals)[name].Move()
	case "speak":
		(*animals)[name].Speak()
	}
}
