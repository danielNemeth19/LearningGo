package main
import (
    "fmt"
)


func test(s string, c chan string){
    fmt.Printf("Sending %s...\n", s)
    c <- s
}


func main(){
    c1 := make(chan string)
    c2 := make(chan string)
    
    go test("first", c1)
    go test("second", c2)
    
/*Select statement:
 *  may have a choice of which data to use (i.e. first-come first-served)
 *  selects waits on the first data from  a set of channels
 */      
    select {
        case a :=<- c1:
            fmt.Printf("...received %s\n", a)
        case b :=<- c2:
            fmt.Printf("...received %s\n", b)
    }
}
