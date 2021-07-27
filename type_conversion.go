package main

import "fmt"


func main() {
    var x int32 = 1
    var y int16 = 2
    
    x = int32(y)
    fmt.Printf("Value of x is: %d\n", x)
}
