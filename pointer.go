package main

import "fmt"


func main() {
    var x int = 1
    var y int
    var ip *int // ip is pointer to int
    
    fmt.Printf("Value of x is:\n")
    fmt.Println(x)
    
    fmt.Printf("Value of y is:\n")
    fmt.Println(y)
    
    fmt.Printf("Value of ip is:\n")
    fmt.Println(ip)
    
    ip = &x
    fmt.Printf("ip now points to x:\n")
    fmt.Println(ip)
    
    y = *ip
    fmt.Printf("y is now 1:\n")
    fmt.Println(y)
    
    ptr:= new(int)
    fmt.Printf("ptr now is a pointer to an int\n")
    fmt.Println(ptr)
    
    *ptr = 3
    fmt.Printf("The memory location ptr is pointing to now stores the value number 3\n")
    fmt.Println(*ptr)
}
