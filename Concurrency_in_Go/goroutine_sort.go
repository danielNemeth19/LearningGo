package main
import (
    "fmt"
    "strconv"
    "sort"
)


func built_in_sort(my_s []int, c chan []int) {
    sort.Ints(my_s)
    c <- my_s
}


func read_in_slice_to_sort() []int {
    var s string
    input_slice := make([]int, 0, 12)
    for {
        fmt.Printf("Enter %v more integers -- or press x! ", 12-len(input_slice))
        fmt.Scanf("%s", &s)
        if s == "x" || s == "X" {
            return input_slice
        }
        number, err := strconv.Atoi(s)
        if err != nil {
            continue
        } else {
            input_slice = append(input_slice, number)
        }
        if len(input_slice) == 12 {
            return input_slice
        }
    }
    return input_slice
}


func main(){
    c := make(chan []int) 
    
    slice_to_sort := read_in_slice_to_sort()
    fmt.Printf("To be sorted: %v -- len of array is: %d\n", slice_to_sort, len(slice_to_sort))
    
    number_of_batches := 4 
    part_size := len(slice_to_sort) / number_of_batches
    modulo := len(slice_to_sort) % number_of_batches
    
    start := 0
    for i:=0; i<number_of_batches; i++ {
        end := start+part_size
        if i == number_of_batches -1 {
            end = end + modulo
        }
        go built_in_sort(slice_to_sort[start:end], c)
        start = start + part_size
    }

    merged_slice := make([]int, 0, len(slice_to_sort))
    first_batch := <- c
    fmt.Printf("Sorted: %v\n", first_batch)
    merged_slice = append(merged_slice, first_batch...)
   
    second_batch := <- c
    fmt.Printf("Sorted: %v\n", second_batch)
    merged_slice = append(merged_slice, second_batch...)
    
    third_batch := <- c
    fmt.Printf("Sorted: %v\n", third_batch)
    merged_slice = append(merged_slice, third_batch...)
    
    fourth_batch := <- c
    fmt.Printf("Sorted: %v\n", fourth_batch)
    merged_slice = append(merged_slice, fourth_batch...)
    sort.Ints(merged_slice)
    fmt.Printf("Merged and sorted: %v\n", merged_slice)
}
