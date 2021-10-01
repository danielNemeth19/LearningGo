package main
import (
    "fmt"
    "time"
)


func send(s string, c chan string, should_wait bool){
    if should_wait{
        time.Sleep(time.Second)
    }
    fmt.Printf("Sending %s...\n", s)
    c <- s
}


func main(){
    c := make(chan string)
    abort := make(chan string)
    
    go send("first", c, false)
    go send("I want to abort", abort, true)

/*Select statement:
 *  use select with a separate abort chanel
 *  may want to receive data until an abort signal received
 */  
for {
    select {
        case a :=<- c:
            fmt.Printf("...received %s\n", a)
        case <- abort:
            fmt.Println("Received abort")
            return
    }
}
}
