package main
import "fmt"


func main(){
    arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
    
    s1 := arr[1:3]
    s2 := arr[2:5]
    
    fmt.Printf("s1 is: %s\n", s1)
    fmt.Printf("s2 is: %s\n", s2)
    
    // length and capacity
    arr2 := [3]string{"y", "g", "s"}
    s3 := arr2[0:1]
    fmt.Printf("arr2 is: %s\n", arr2)
    fmt.Printf("s3 is: %s\n", s3)
    fmt.Printf("Len of arr2 is: %d - capaciy of arr2 is: %d\n", len(s3), cap(s3))
    
    // Accessing Slices
    fmt.Printf("s1 second element is is: %s\n", s1[1])
    fmt.Printf("s2 first element is: %s\n", s2[0])
    
    // Slice literals -> creates underlying array and creates a slice to reference it
    // Slice points to the start of the array, length is capacity
    my_slice := []int{1, 2, 3}
    fmt.Printf("Len of my_slice is: %d - capaciy of my_slice is: %d\n", len(my_slice), cap(my_slice))
    
    // Writing to a slice changes underlying array - but not other array referencing that array
    arr3 := [2]string{"Hello", "World"}
    
    s4 := arr3[0:2]
    s5 := arr3[0]
    
    fmt.Printf("s4 is: %s\n", s4)
    fmt.Printf("s5 is: %s\n", s5)
    s4[0] = "Bye"
    fmt.Printf("s4 is: %s\n", s4)
    fmt.Printf("s5 is: %s\n", s5)
    
    fmt.Printf("arr3 is: %s\n", arr3)
    
}
