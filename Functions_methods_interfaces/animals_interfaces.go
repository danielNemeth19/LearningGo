package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "reflect"
)


type Animal interface {
    Eat()
    Move()
    Speak()
}


type Cow struct {
    name string
}


func (c Cow) Eat() {
    food := "grass"
    fmt.Printf("Food eaten by %s is %s\n", c.name, food)
}


func (c Cow) Move() {
    locomtion := "walk"
    fmt.Printf("Locomtion of %s is %s\n", c.name, locomtion)
}


func (c Cow) Speak() {
    sound := "moo"
    fmt.Printf("Sound of %s is %s\n", c.name, sound)
}


type Bird struct {
    name string
}


func (b Bird) Eat() {
    food := "worms"
    fmt.Printf("Food eaten by %s is %s\n", b.name, food)
}


func (b Bird) Move() {
    locomtion := "fly"
    fmt.Printf("Locomtion of %s is %s\n", b.name, locomtion)
}


func (b Bird) Speak() {
    sound := "peep"
    fmt.Printf("Sound of %s is %s\n", b.name, sound)
}


type Snake struct {
    name string
}


func (s Snake) Eat() {
    food := "mice"
    fmt.Printf("Food eaten by %s is %s\n", s.name, food)
}


func (s Snake) Move() {
    locomtion := "slither"
    fmt.Printf("Locomtion of %s is %s\n", s.name, locomtion)
}


func (s Snake) Speak() {
    sound := "hsss"
    fmt.Printf("Sound of %s is %s\n", s.name, sound)
}


func _validate_command(input_s string) bool {
    switch input_s {
        case 
            "newanimal",
            "query":
            return true
    }
    return false
}


func _validate_subject(s1, s2 string) bool {
    if s1 == "newanimal" {
        switch s2 {
            case 
                "cow",
                "bird",
                "snake":
                return true
        }
    }
    if s1 == "query" {
        switch s2 {
            case
                "eat",
                "move",
                "speak":
                return true
        }
    }
    return false
}


func parse_inputs() []string {
    fmt.Printf("> ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    inputString := scanner.Text()
    inputSlice := strings.Split(inputString, " ")
    return inputSlice
}


func store_animal(name, subj string, animals map[string]Animal) bool {
    _, flag := animals[name]
    if flag {
        fmt.Printf("Animal named %s exists already!\n", name)
        return false
    }
    switch subj {
        case "cow":
            animals[name] = Cow{name}
        case "bird":
            animals[name] = Bird{name}
        case "snake":
            animals[name] = Snake{name} 
    }
    fmt.Println("Created it!")
    return true
}


func query_animal(name, subj string, animals map[string]Animal) bool {
    animal, flag := animals[name]
    if !flag {
        fmt.Printf("Animal with name %s doesn't exist!\n", name)
        return false
    }
    meth := reflect.ValueOf(animal).MethodByName(strings.Title(subj))
    meth.Call(nil)
    return true
}


func main(){
    animals := make(map[string]Animal)
    for {
        input := parse_inputs()
        if len(input) != 3 {
            fmt.Printf("Line should have three parts (command animal subj): got %d\n", len(input))
            continue
        }
        command, name, subj := input[0], input[1], input[2]        
        if !_validate_command(command){
            fmt.Printf("Input given as command not recognized: %s (Should be one of these: query, newanimal)\n", command)
            continue
        }
        if !_validate_subject(command, subj){
            fmt.Printf("Input given as subject of command not recognized: %s\n", subj)
            continue
        }
        if command == "newanimal" {
            stored := store_animal(name, subj, animals)
            if !stored {
                continue
            }
        } else {
            exists := query_animal(name, subj, animals)
            if !exists {
                continue 
            }
        }
    }
}
