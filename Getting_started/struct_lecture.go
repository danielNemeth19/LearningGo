package main
import "fmt"


type Person struct  {
    name string
    addr string
    phone string
}

func main(){
    var p1 Person
    fmt.Printf("Type of p1 variable is: %T\n", p1)
    
    p1.name = "Daniel"
    p1.addr = "Budapest"
    
    fmt.Printf("Address of: %s is %s\n", p1.name, p1.addr)
    
    // stuct can be also initialized with 'new'
    p2 := new(Person)
    fmt.Printf("If struct initialized with empty strings -> Address of: %s is %s\n", p2.name, p2.addr)
    
    // struct can also be initialized using struct literal
    p3 := Person{name: "Reka", addr:"Budapest", phone: "+3670258"}
    fmt.Printf("Address of: %s is %s\n", p3.name, p3.addr)
    
    
}
