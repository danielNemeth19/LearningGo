package main
import (
    "fmt"
)


/* Function are First-class
Functions can be treated like other types - variables can be˙declared with function type
    - can be created dynamically
    - can be passed as arguments and returned as values
    - can be stored in data structures
*/

// Declaring a variable as a function
var funcVar func(int) int
    func incFn(x int) int {
        return x + 1
    }
    
    
func incFn2(x int) int {
    return x + 2
}

// Function as argument
func applyIt(afunct func (int) int, val int) int {
    return afunct(val)
}


func decFn(x int) int {
    return x - 1
}


func main(){
    funcVar = incFn
    fmt.Println(funcVar(1))
    
    test := incFn2
    fmt.Println(test(4))
    
    fmt.Println(applyIt(decFn, 2))
    
    // creating anonymous function
    v := applyIt(func (x int) int {return x + 1}, 2)
    fmt.Println(v)
}
