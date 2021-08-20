package main
import (
    "fmt"
    "LearningGo/Functions_methods_interfaces/data"
)


func main(){
    data.PrintX()
    
    var p data.Point
    p.InitMe(2,4)
    p.PrintMe()
    p.Scale(3)
    p.PrintMe()
    
/*
* OffsetX has its argument defined as a pointer type (pointer receiver)
* But no need to reference when calling the method:
*   - eventhough p is the actual struct, no need to write &p.OffsetX
*   - this is done automatically by the compiler (syntactic sugar) 
*/
    p.OffsetX(10)
    p.PrintMe()
    
    fmt.Println("Bad method definition (call by value) won't change x's value:")
    p.OffsetXBad(10)
    p.PrintMe()
}
