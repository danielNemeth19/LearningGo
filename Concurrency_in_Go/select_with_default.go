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
    c1 := make(chan string)
    c2 := make(chan string)
    
    go send("first", c1, false)
    go send("second", c2, false)

/*Default select:
 *  may want a default operation to avoid blocking
 */  

select {
    case a :=<- c1:
        fmt.Printf("...received %s\n", a)
    case b:=<- c2:
        fmt.Printf("...received %s\n", b)
    default:
        fmt.Println("nop")
    }
}
