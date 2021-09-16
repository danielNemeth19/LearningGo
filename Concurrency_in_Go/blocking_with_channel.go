package main
import (
    "fmt"
)


func foo(c chan string){
    fmt.Printf("New Routine\n")
    c <- "I could send anything --- it will be thrown away"
}

/*
 * Blocking and Synchronization
 * Channel communication is synchronous
 *  - blocking is the same as waiting for communication
 *  - receiving and ignoring the result is same as a Wait() using WaitGroup
 */

func main(){
    c := make(chan string)
    
    go foo(c)
    <- c
    fmt.Printf("Main Routine\n")
}
