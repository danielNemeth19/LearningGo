package main
import "fmt"


func foo() *int {
    x:= 1
    return &x
}


func main() {
    var y *int
    y = foo()
    fmt.Printf("Address of x is: %d\n", y)
    fmt.Printf("Value of x is: %d\n", *y)
    fmt.Printf("Address of y is: %d\n", &y)
}
