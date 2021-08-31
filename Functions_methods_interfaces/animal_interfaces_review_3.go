package main

import "fmt"

var (
	COW_TYPE         = "cow"
	COW_FOOD         = "grass"
	COW_LOCOMOTION   = "walk"
	COW_SOUND        = "moo"
	BIRD_TYPE        = "bird"
	BIRD_FOOD        = "worms"
	BIRD_LOCOMOTION  = "fly"
	BIRD_SOUND       = "peep"
	SNAKE_TYPE       = "snake"
	SNAKE_FOOD       = "mice"
	SNAKE_LOCOMOTION = "slither"
	SNAKE_SOUND      = "hsss"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

func (ap Cow) Eat() {
	fmt.Printf("%v eat %v.\n", ap.Name, ap.Food)
}

func (ap Cow) Move() {
	fmt.Printf("%v move by %v.\n", ap.Name, ap.Locomotion)
}

func (ap Cow) Speak() {
	fmt.Printf("%v sound like %v.\n", ap.Name, ap.Sound)
}

type Bird struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

func (ap Bird) Eat() {
	fmt.Printf("%v eat %v.\n", ap.Name, ap.Food)
}

func (ap Bird) Move() {
	fmt.Printf("%v move by %v.\n", ap.Name, ap.Locomotion)
}

func (ap Bird) Speak() {
	fmt.Printf("%v sound like %v.\n", ap.Name, ap.Sound)
}

type Snake struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

func (ap Snake) Eat() {
	fmt.Printf("%v eat %v.\n", ap.Name, ap.Food)
}

func (ap Snake) Move() {
	fmt.Printf("%v move by %v.\n", ap.Name, ap.Locomotion)
}

func (ap Snake) Speak() {
	fmt.Printf("%v sound like %v.\n", ap.Name, ap.Sound)
}

func main() {

	var animalQuery string
	var animalName string
	var animalInterface string
	animalMap := make(map[string]Animal)

	for animalQuery != "X" && animalName != "X" && animalInterface != "X" {
		isCorrect := true
		fmt.Println("\nPlease enter your command [newanimal animalname animaltype] or [query animalname animalactivity].")
		fmt.Println("The valid animaltype are cow, bird or snake.")
		fmt.Println("The valid animalactivity are eat, move or speak.")
		fmt.Println("Enter 'X X X' to exit")
		fmt.Print("> ")
		fmt.Scan(&animalQuery, &animalName, &animalInterface)
		if (animalQuery == "X") && (animalName == "X") && (animalInterface == "X") {
			fmt.Println("Bye Bye!!")
		} else {
			fmt.Printf("> Your input data are [%v %v %v].\n", animalQuery, animalName, animalInterface)

			switch animalQuery {
			case "newanimal":
				if animalInterface != "cow" && animalInterface != "bird" && animalInterface != "snake" {
					isCorrect = false
					fmt.Println("!!! Invalid animaltype you can enter only 'cow', 'bird' or 'snake' !!!")
				}
			case "query":
				if animalInterface != "eat" && animalInterface != "move" && animalInterface != "speak" {
					isCorrect = false
					fmt.Println("!!! Invalid animalactivity you can enter only 'eat', 'move' or 'speak' !!!")
				}
			default:
				isCorrect = false
				fmt.Println("!!! Invalid command you can enter only 'newanimal' or 'query' !!!")
			}

			if isCorrect {
				switch animalQuery {
				case "newanimal":

					// Check duplicate name
					_, ok := animalMap[animalName]
					if ok {
						// Duplicate name
						isCorrect = false
						fmt.Println("!!! Cannot create your animal. The name is duplicated !!!")
					} else {
						// Create animal
						switch animalInterface {
						case "cow":
							animalMap[animalName] = Cow{Name: animalName, Food: COW_FOOD, Locomotion: COW_LOCOMOTION, Sound: COW_SOUND}
						case "bird":
							animalMap[animalName] = Bird{Name: animalName, Food: BIRD_FOOD, Locomotion: BIRD_LOCOMOTION, Sound: BIRD_SOUND}
						case "snake":
							animalMap[animalName] = Snake{Name: animalName, Food: SNAKE_FOOD, Locomotion: SNAKE_LOCOMOTION, Sound: SNAKE_SOUND}
						default:
							isCorrect = false
							fmt.Println("!!! Something is error. !!!")
						}
						if isCorrect {
							fmt.Println("Created it!")
						}
					}
				case "query":
					// Check exists name
					animal, ok := animalMap[animalName]
					if !ok {
						// not exists name
						isCorrect = false
						fmt.Printf("!!! Your animal name {%v} doesn't exits !!!\n", animalName)
					} else {
						// Create animal
						switch animalInterface {
						case "eat":
							animal.Eat()
						case "move":
							animal.Move()
						case "speak":
							animal.Speak()
						default:
							isCorrect = false
							fmt.Println("!!! Something is error. !!!")
						}
					}
				default:
					isCorrect = false
					fmt.Println("!!! Something is error. !!!")
				}
			}
		}
	} // end for
}
