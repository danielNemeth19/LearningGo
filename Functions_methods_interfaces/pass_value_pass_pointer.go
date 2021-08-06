package main
import (
    "fmt"
)


type Person struct {
    firstName string
    lastName string
}


func changeNameValue(p Person) {
    p.firstName = "BobValue"
    fmt.Printf("From function changeNameValue: %s\n", p)
}


func changeNamePointer(p *Person) {
    p.firstName = "BobPointer"
    fmt.Printf("From function changeNamePointer: %s\n", p)
}


func main(){
    person := Person{
        firstName: "Elek",
        lastName: "Mekk",
    }
    
    fmt.Printf("Before calling any function: %s\n", person)
    changeNameValue(person)
    fmt.Printf("After calling changeNameValue: %s\n", person)
    
    changeNamePointer(&person)
    fmt.Printf("After calling changeNamePointer: %s\n", person)
    
}
