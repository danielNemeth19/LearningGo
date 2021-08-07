package main
import (
    "fmt"
)


type Person struct {
    firstName string
    lastName string
}

// pass by value won't change the calling environment
// passed arguments are copied to parameters
// function operates on a copy of the value of the passed in argument
func changeNameValue(p Person) {
    p.firstName = "BobValue"
    fmt.Printf("From function changeNameValue: %s\n", p)
}


// passed in argument can be a pointer
// this case function has direct access to caller variable in memory - and it can change it
func changeNamePointer(p *Person) {
    p.firstName = "BobPointer"
    fmt.Printf("From function changeNamePointer: %s\n", *p)
}


// if value is returned it can be saved into new variable
func changeNameValueReturn(p Person) Person {
    p.firstName = "BobValueReturn"
    fmt.Printf("From function changeNameValueReturn: %s\n", p)
    return p
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
    
    person_2 := changeNameValueReturn(Person{firstName: "Elek", lastName: "Mekk"})
    fmt.Printf("After calling changeNameValueReturn: %s\n", person_2)
    
}
