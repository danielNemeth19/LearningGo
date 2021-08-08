package main
import (
    "fmt"
)


/*
Call can be deferred until surrounding function completes
Typically used for cleanup activites  
*/
func example_1(){
    defer fmt.Println("Bye")
    fmt.Println("Hi from example_1!")
}


/*
Arguments of a deferred call are evaluated immediately
This means only the call is deferred!!
*/
func example_2(){
    i := 1
    defer fmt.Printf("i from example_2: %v\n", i+1)
    i++
    fmt.Println("Hello from example_2")
}


func main(){
    example_1()
    example_2()
}
