package main
import "fmt"

func main(){
    type Grades int
    const(
        A Grades = iota
        B
        C
        D
        F
    )
    
    fmt.Printf("The value of A is %v\n", A)
    fmt.Printf("The value of B is %v\n", B)
    fmt.Printf("The value of C is %v\n", C)
    fmt.Printf("The value of D is %v\n", D)
}
