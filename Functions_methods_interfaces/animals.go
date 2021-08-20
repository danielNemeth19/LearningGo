package main
import (
    "fmt"
    "bufio"
    "strings"
    "os"
    "reflect"
)


type Animal struct {
    food string
    locomotion string
    noise string
}


func (a Animal) Eat(){
    fmt.Printf("%s\n", a.food)
}


func (a Animal) Move(){
    fmt.Printf("%s\n", a.locomotion)
}


func (a Animal) Speak(){
    fmt.Printf("%s\n", a.noise)
}


func get_struct_map() map[string]Animal {
    cow := Animal{food:"grass", locomotion:"walk", noise:"moo"}
    bird := Animal{food:"worms", locomotion:"fly", noise:"peep"}
    snake := Animal{food:"mice", locomotion:"slither", noise:"hss"}
    animal_map := map[string]Animal {
        "cow": cow,
        "bird": bird,
        "snake": snake,
    }
    return animal_map
}


func _validate_animal(input_s string) bool {
    switch input_s {
        case 
            "cow",
            "bird",
            "snake":
            return true
    }
    return false
}


func _validate_query(input_s string) bool {
    switch input_s {
        case 
            "eat",
            "move",
            "speak":
            return true
    }
    return false
}


func parse_inputs() (string, string) {
    fmt.Printf("> ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    inputString := scanner.Text()
    inputSlice := strings.Split(inputString, " ")
    if len(inputSlice) != 2 {
        fmt.Printf("Line should have two parts (animal query): got %d\n", len(inputSlice))
        os.Exit(1)
    }
    if !_validate_animal(inputSlice[0]){
        fmt.Printf("Input given as animal not recognized: %s (Should be one of these: cow, bird, snake)\n", inputSlice[0])
        os.Exit(1)
    }
    if !_validate_query(inputSlice[1]){
        fmt.Printf("Input given as query not recognized: %s (Should be one of these: eat, move, speak)\n", inputSlice[1])
        os.Exit(1)
    }
    return inputSlice[0], inputSlice[1]
}


func main() {
    animal_map := get_struct_map()
    
    for {
        animal, query := parse_inputs()    
        animal_struct := animal_map[animal]
        meth := reflect.ValueOf(animal_struct).MethodByName(strings.Title(query))
        meth.Call(nil)
    }
}
