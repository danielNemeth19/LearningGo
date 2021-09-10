// PACKAGE MAIN  ===============================================
package main

// IMPORT  ===============================================
import(
  "fmt"
  "strings"
  "bufio"
  "os"
  "reflect"
)


// INTERFACE DECLARATION ===============================================
type Animal interface{
  Eat() //method should print the animal’s food.
  Move() //method should print the animal’s locomotion.
  Speak() //method should print the animal’s spoken sound.
}


// TYPE DECLARATION ===============================================

// COW TYPE ///////////////////////////////////////////////////////
type Cow struct{
  Name string

}

// BIRD TYPE ///////////////////////////////////////////////////////
type Bird struct{
  Name string

}

// SNAKE TYPE /////////////////////////////////////////////////////
type Snake struct{
  Name string

}

// STRUCT ANIMAL STORAGE WITH  SLICES (COW, BIRD AND SNAKE) TO STORE THE DATA
type AnimalsStorage struct{
  CowAnimals []*Cow
  BirdAnimals []*Bird
  SnakeAnimals []*Snake

}

// STRUCT ID TO COMPOSE A MAP(NAME ANIMAL:ID) IN ORDER TO FASTER THE QUERY FUNCTION
type ID struct{
  SliceString string
  AnimalIndex int
}
// METHODS DECLARATION =================================================

// COW //////////////////////////////////////////////////////////////

// Eat()
func (c Cow) Eat() {
    fmt.Printf("The <<< %s >>> is a type of Cow that eats <<< grass >>>\n", c.Name)

}
// Move()
func (c Cow) Move() {
    fmt.Printf("The <<< %s >>> is a type of Cow that moves by <<< walk >>>\n", c.Name)
}
// Speak()

func (c Cow) Speak() {
    fmt.Printf("The <<< %s >>> is a type of Cow that speaks with <<< moo >>>\n", c.Name)
}


// BIRD //////////////////////////////////////////////////////////////

// Eat()
func (b Bird) Eat() {
  fmt.Printf("The <<< %s >>> is a type of Bird that eats <<< worms >>>\n", b.Name)

}
// Move()
func (b Bird) Move() {
  fmt.Printf("The <<< %s >>> is a type of Bird that moves by <<< fly >>>\n", b.Name)
}
// Speak()

func (b Bird) Speak() {
  fmt.Printf("The <<< %s >>> is a type of Bird that speaks with <<< peep >>>\n", b.Name)
}


// SNAKE ////////////////////////////////////////////////////////////////

// Eat()
func (s Snake) Eat() {
    fmt.Printf("The <<< %s >>> is a type of Snake that eats <<< mice >>>\n", s.Name)

}
// Move()
func (s Snake) Move() {
    fmt.Printf("The <<< %s >>> is a type of Snake that moves by <<< slither >>>\n", s.Name)
}
// Speak()

func (s Snake) Speak() {
    fmt.Printf("The <<< %s >>> is a type of Snake that speaks with <<< hsss >>>\n", s.Name)
}

// VARIABLES DECLARATION ===================================================
var storage = new(AnimalsStorage)
var storageMap = make(map[string]ID)
var command_01, command_02, command_03 string



// FUNCTIONS DECLARATION ====================================================

// Presentation ///////////////////////////////////////////////////////////
func presentation(){
    fmt.Println(strings.Repeat("=", 70),"\n")
    fmt.Println("##### Create or Query Animals from Cow, Bird and Snake types #####")
    fmt.Println("\n",strings.Repeat("=", 70),"\n")
    fmt.Println("INSTRUCTIONS:\n1) Wait the cursor ((( >>>: )))\n\n")
    fmt.Println("2)### Create New Animal: ###\na)Command newanimal + \nb)Name or Breed of your choice + \nc)Choose one of 3 types: Cow, Bird and Snake \n   2.1) Type 3 commands separated by space: newanimal | name | type \n   2.2) E.g: newanimal parrot bird\n   2.3) Check confirmation message\n\n")
    fmt.Println("3)### Query Animal:  ###\na)Command query + \nb)Name or Breed  + \nc)Choose one of 3 informations: eat, move or speak \n   3.1) Type 3 commands separated by space: query | name | information \n   3.2) E.g: query parrot eat\n   3.3) Check confirmation message\n\n")
    fmt.Println("\n",strings.Repeat("=", 70),"\n")
}


// Get User input ////////////////////////////////////////////////////////
func getUserInput(){
    box := bufio.NewScanner(os.Stdin)
    box.Scan()
    input := strings.Fields(box.Text())

    if input[0] == "exit" || input[0] == "EXIT" || input[0] == "Exit" {
        fmt.Println("\nExiting...")
        os.Exit(0)

    }   else if input[0] == "query" || input[0] == "newanimal" {
                switch {
                  case len(input) < 4:
                      command_01 = strings.ToLower(input[0])
                      command_02 = strings.ToLower(input[1])
                      command_03 = strings.ToLower(input[2])
                  default:
                    fmt.Println("Error: Input can have only a command of 3 words per time, please check the instructions ")
                    os.Exit(0)
                }

        }    else {
              fmt.Println("Error: Invalid Command, please check the instructions")
              os.Exit(0)
             }

}

// Store Animals ////////////////////////////////////////////////////////

func storeAnimals(a Animal) bool{

  var done = false
  switch RealType := a.(type) {
    case *Cow:
      storage.CowAnimals = append(storage.CowAnimals, RealType)
      index := len(storage.CowAnimals)-1
      nameAnimal := storage.CowAnimals[index].Name
      //id := new(ID)
      var id ID
      id.SliceString = "CowAnimals"
      id.AnimalIndex = index
      storageMap[nameAnimal] = id
      fmt.Println("\n### The New Animal is now stored\n\nanimal ===> ", nameAnimal )
      done = true

    case *Bird:
      storage.BirdAnimals = append(storage.BirdAnimals, RealType)
      index := len(storage.BirdAnimals)-1
      nameAnimal := storage.BirdAnimals[index].Name
      //id := new(ID)
      var id ID
      id.SliceString = "BirdAnimals"
      id.AnimalIndex = index
      storageMap[nameAnimal] = id
      fmt.Println("\n### The New Animal is now stored\n\nanimal ===> ", nameAnimal )
      done = true

    case *Snake:
      storage.SnakeAnimals = append(storage.SnakeAnimals, RealType)
      index := len(storage.SnakeAnimals)-1
      nameAnimal := storage.SnakeAnimals[index].Name
      //id := new(ID)
      var id ID
      id.SliceString = "SnakeAnimals"
      id.AnimalIndex = index
      storageMap[nameAnimal] = id
      fmt.Println("\n### The New Animal is now stored\n\nanimal ===> ", nameAnimal )
      done = true

    default:
      fmt.Println(RealType)
      fmt.Println(reflect.TypeOf(RealType))
      fmt.Println("\n### The New Animal Type does not exist\n")
      os.Exit(0)
  }

    return done
}


/*
Each “newanimal” command must be a single line containing three strings.
The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal.
The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.
*/

func newanimal(nameNewAnimal, typeAnimal string) {
  // commad: newanimal | name | type
    _, ok := storageMap[nameNewAnimal]
    if ok {
          fmt.Printf("Error: %s already exists. Choose another name and try again!\n", nameNewAnimal)
          os.Exit(0)

    } else {
            if typeAnimal == "cow" || typeAnimal == "bird" || typeAnimal == "snake" {
                switch typeAnimal {
                  case "cow":
                      animal := new(Cow)
                      animal.Name = nameNewAnimal
                      storeAnimals(animal)
                      fmt.Println("Created it!")
                  case "bird":
                      animal := new(Bird)
                      animal.Name = nameNewAnimal
                      storeAnimals(animal)
                      fmt.Println("Created it!")
                  case "snake":
                      animal := new(Snake)
                      animal.Name = nameNewAnimal
                      fmt.Println(reflect.TypeOf(animal))
                      storeAnimals(animal)
                      fmt.Println("Created it!")
                }
            } else {
                fmt.Printf("Error: Invalid Type. Only: Cow, Bird and Snake are allowed! Please read the instructions and try again")
                os.Exit(0)
              }
       }

}

/*
Each “query” command must be a single line containing 3 strings.
The first string is “query”.
The second string33 is the name of the animal.
The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.
*/

func query(nameAnimal, informationAnimal string) {
  // commad: query | name | information (“eat”, “move”, or “speak”)
  value, ok := storageMap[nameAnimal]
  if ok {
      sliceAnimalStored := value.SliceString
      id_position := value.AnimalIndex

      if sliceAnimalStored == "CowAnimals"{
            cowTemp := storage.CowAnimals[id_position]
            switch informationAnimal {
                case "eat":
                    cowTemp.Eat()
                case "move":
                    cowTemp.Move()
                case "speak":
                    cowTemp.Speak()
                default:
                    fmt.Printf("The %s is not a valid information request",informationAnimal)
                    os.Exit(0)
            }


      }  else if sliceAnimalStored == "BirdAnimals"{
                  birdTemp := storage.BirdAnimals[id_position]
                  switch informationAnimal {
                      case "eat":
                          birdTemp.Eat()
                      case "move":
                          birdTemp.Move()
                      case "speak":
                          birdTemp.Speak()
                      default:
                          fmt.Printf("The %s is not a valid information request",informationAnimal)
                          os.Exit(0)
                  }

         }  else if sliceAnimalStored == "SnakeAnimals"{
                        snakeTemp := storage.SnakeAnimals[id_position]
                        switch informationAnimal {
                            case "eat":
                                snakeTemp.Eat()
                            case "move":
                                snakeTemp.Move()
                            case "speak":
                                snakeTemp.Speak()
                            default:
                                fmt.Printf("The %s is not a valid information request",informationAnimal)
                                os.Exit(0)

                         }
            }

  } else{
       fmt.Println("Error: This animal does not exist, please create it first ")
       os.Exit(0)
    }


}



// MAIN FUNCTION ==============================================================
func main (){
    presentation()
    for {

        fmt.Printf("\n### Type a 3 words command or Exit: ###\n>>>: ")
        getUserInput()
        switch command_01 {
          case "newanimal":
              newanimal(command_02, command_03)
          case "query":
              query(command_02, command_03)
          default:
              fmt.Println("Error: ### Invalid Command ###\nPlease check the instruction and try again\n")
              os.Exit(0)
        }
    }
}
