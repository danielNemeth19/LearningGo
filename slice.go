package main
import (
    "fmt"
    "strings"
    "sort"
    "os"
    "bufio"
    "strconv"
)


func _convert(s string) string {
    return strings.Trim(s, "\n")
}


func main(){
    
    my_slice := make([]int, 0, 3)
    for true {
        fmt.Println("Enter an integer!")
        in := bufio.NewReader(os.Stdin)
        inputSting, _ := in.ReadString('\n')
        s := _convert(inputSting)
        
        if s == "X" {
            fmt.Println("Exit signal received")
            os.Exit(1)
        }
        
        number, err := strconv.Atoi(s)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
        }
        if err == nil {
            my_slice = append(my_slice, number)
            sort.Ints(my_slice)
            fmt.Println(my_slice) 
        }
    }
}
