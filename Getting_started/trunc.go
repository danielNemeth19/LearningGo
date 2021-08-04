package main
import (
    "fmt"
)


func main(){
    var myFloat float64
    
    fmt.Printf("Enter a floating point number:\n")
    _, err := fmt.Scan(&myFloat)
    if err != nil {
        fmt.Printf("%T, %v\n", err, err)
    }
    fmt.Printf("Truncated number: %d\n", int(myFloat))
    
}
