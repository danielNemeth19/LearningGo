package main
import (
    "fmt"
)


func prod(v1 int, v2 int, c chan int){
    c <- v1 * v2
}

/*
 * Unbuffered channel
 *  - this is the default channel behaviour
 *  - cannot hold data in transit
 *  - sending blocks until data is received
 *  - receiving blocks until data is sent
 *  
*/

func main(){
    c := make(chan int)
    go prod(1, 2, c)
    go prod(3, 4, c)
        
    a := <-c
    b := <-c
    
    fmt.Printf("Result of a: %d\n", a)
    fmt.Printf("Result of b: %d\n", b)
    
    fmt.Printf("Final result: %d\n", a*b)
}
