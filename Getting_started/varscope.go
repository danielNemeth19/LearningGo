package main

import "fmt"

var x int = 3

func f(){
    fmt.Printf("Value of x is: %d\n", x)
}

func g(){
    fmt.Printf("Value of x is: %d\n", x)
}

func h(){
    var y int = 5
    fmt.Printf("Value of y is: %d\n", y)
}

/* With below function program would not compile
func k(){
     fmt.Printf("Value of y is: %d\n", y)
}
*/


func main() {
    f()
    g()
    h()
}
