package main
import (
    "fmt"
)


func send(s string, c chan string){
    fmt.Printf("Sending %s...\n", s)
    c <- s
}


func read(c chan string){
    b :=<- c
    fmt.Printf("Received: %s\n", b)
}


func main(){
    inchan := make(chan string)
    outchan := make(chan string)
    
    go send("first", inchan)
    go read(outchan)

/*Select statement:
 *  select is possible even between send and receive operations
 *  reading from inchan blocks until no data has been sent; sending to outchan blocks until nobody is receiving on outchan
 *  whichever operation is unblocked first -> that's executed
 *  - to demonstrate that 'case a' can be forced, import time and wait in the read function for a second - time.Sleep(time.Second)
 */  
    select {
        case a :=<- inchan:
            fmt.Printf("...received %s\n", a)
        case outchan <- "second":
            fmt.Printf("sent data to outchan\n")
    }
}
