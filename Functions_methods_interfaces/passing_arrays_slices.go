package main
import (
    "fmt"
)


// array arguments are copied
// this is problematic for big arrays
func passing_array_as_value(x [3]int) int {
    return x[0]
}


// it is possible to pass array pointer
// this is messy and unnecessary 
func passing_array_pointer(x *[3]int) {
     (*x)[0] = (*x)[0] + 1    
}


/* slices contain a pointer to the underlying array!!
passing a slice copies the pointer: a slice really is a structure that contain three things:
    - a pointer to the start to the slice in the array
    - a length
    - and a capacity
So this is call by value, but since we're passing a slice, what we copy is a pointer! Hence no return needed to modify the calling environment
*/
func pass_slice(sli []int) {
    sli[0] = sli[0] + 1
}


func main(){
    a := [3]int{1, 2, 3}
    fmt.Printf("Before calling any function: %d\n", a)
    
    fmt.Printf("Passing array as value: %d\n", passing_array_as_value(a))
    
    passing_array_pointer(&a)
    fmt.Printf("After calling passing_array_pointer: %d\n", a)
    
    b := []int{1, 2, 3}
    pass_slice(b)
    fmt.Printf("After calling pass_slice: %d\n", b)
}
