package main
import (
    "fmt"
)


type Myint int

/* Receiver type defined before function name:
    - Myint is the reveiver type
    - mi is the variable that refers to the particular receiver object that this function will be called on
    - note that function operates on a copy of the object and returns the copy (it won't modify it)
 */
func (mi Myint) Double() int {
    return int(mi*2)
}



/* 
 * Use dot notation to call the method
 * Note implicit method argument:
 *  - Double has no argument(s) defined
 *  - but the object itself ('v' in this case) is passed to the function!
 *  - this a call by value, so a copy of v is created and passed to Double when it's called
 */
func calling_receiver_type() {
    v := Myint(3)
    fmt.Printf("Value of v: %d\n", v)
    fmt.Printf("Calling v's Double method: %d\n", v.Double())
    fmt.Printf("Value of v: %d\n", v)
}


func main(){
    calling_receiver_type()
}

