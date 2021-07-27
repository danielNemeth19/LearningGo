package main

import "fmt"


func main(){
    // create a slice (and array) using make()
    // 2-argument version: specify type and length/capacity
    // initialize to zero, length = capacity
    slice_1 := make([]int, 10)
    fmt.Printf("Len of slice_1 is: %d - capaciy of slice_1 is: %d\n", len(slice_1), cap(slice_1))
    
    // 3-argument version: specify length and capacity separately
    slice_2 := make([]int, 5, 10)
    fmt.Printf("Len of slice_2 is: %d - capaciy of slice_2 is: %d\n", len(slice_2), cap(slice_2))
    
    // Size of a slice can be increased by append()
    // Adds elements to the end of a slice
    // Inserts into underlying array
    // Increases size of array if necessary
    slice_3 := make([]int, 0, 3)
    fmt.Printf("Before append: len of slice_3 is: %d - capaciy of slice_3 is: %d\n", len(slice_3), cap(slice_3))
    slice_3 = append(slice_3, 100)
    fmt.Printf("After append: len of slice_3 is: %d - capaciy of slice_3 is: %d\n", len(slice_3), cap(slice_3))
    fmt.Printf("First element of slice_3: %d\n", slice_3[0])
    
    
}
