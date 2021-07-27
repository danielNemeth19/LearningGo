package main
import "fmt"


func main(){
    my_slice := make([]int, 0, 3)
    fmt.Println(my_slice)
    fmt.Printf("Before append: len of my_slice is: %d - capaciy of my_slice is: %d\n", len(my_slice), cap(my_slice))
    
    my_slice = append(my_slice, 1)
    fmt.Println(my_slice)
    fmt.Printf("After append: len of my_slice is: %d - capaciy of my_slice is: %d\n", len(my_slice), cap(my_slice))
    fmt.Printf("First element of my_slice: %d\n", my_slice[0])
    
    my_slice = append(my_slice, 2)
    fmt.Printf("After append: len %d - capaciy: %d\n", len(my_slice), cap(my_slice))
    
    my_slice = append(my_slice, 3)
    fmt.Printf("After append: len %d - capaciy: %d\n", len(my_slice), cap(my_slice))
    
    my_slice = append(my_slice, 4)
    fmt.Printf("After append: len %d - capaciy: %d\n", len(my_slice), cap(my_slice))
    
}
