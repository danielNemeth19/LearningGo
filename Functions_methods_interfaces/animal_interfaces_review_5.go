package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

// Define interface
type Animal interface {
  Eat()
  Move()
  Speak()
}

// Define types
type Cow struct {}

type Bird struct {}

type Snake struct {}

// Define functions
func (c Cow) Eat() {
  fmt.Printf("grass")
}

func (c Cow) Move() {
  fmt.Printf("walk")
}

func (c Cow) Speak() {
  fmt.Printf("moo")
}

func (b Bird) Eat() {
  fmt.Printf("worms")
}

func (b Bird) Move() {
  fmt.Printf("fly")
}

func (b Bird) Speak() {
  fmt.Printf("peep")
}

func (s Snake) Eat() {
  fmt.Printf("mice")
}

func (s Snake) Move() {
  fmt.Printf("slither")
}

func (s Snake) Speak() {
  fmt.Printf("hsss")
}


func main () {
  // Hard-code data
  db := make(map[string]Animal)

  // Create reader (accepts a sequence separated by white spaces)
  reader := bufio.NewReader(os.Stdin)

  // Prompt
  fmt.Printf("Enter your newanimal or query")

  for {
    fmt.Printf("\n> ")
    input, err := reader.ReadString('\n')
    if err != nil {
      fmt.Printf("An error occured!")
      break
    } else {
      params := strings.Fields(input)

      switch params[0] {
        case "newanimal":
          switch params[2] {
            case "cow":
              var a Cow
              db[params[1]] = a
            case "bird":
              var a Bird
              db[params[1]] = a
            case "snake":
              var a Snake
              db[params[1]] = a
          }
          fmt.Printf("db: %v", db)
          fmt.Printf("Created it!")
        case "query":
          switch params[2] {
            case "eat":
              a := db[params[1]]
              a.Eat()
            case "move":
              a := db[params[1]]
              a.Move()
            case "speak":
              a := db[params[1]]
              a.Speak()
          }
      }
    }
  }
}
