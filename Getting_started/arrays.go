package main
import "fmt"

func basic_array(){
    // Indices starts at 0, elements are initialized to 0
	var x [5]int
	x[0] = 2
	fmt.Printf("Second element of x is: %d\n", x[1])
    
    // Array can be defined as array literal
    var y [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Printf("Fifth element of y is: %d\n", y[4])
    
    z := [...]int{1, 2, 3, 4}
    fmt.Printf("%d\n", z[2])
    
    q := [3]int{10, 11, 12}
    fmt.Printf("%d\n", q[2])
    for i, v := range q {
        fmt.Printf("Index: %d, value is: %d\n", i, v)
    }
    
}

func main(){
	basic_array()
}
