package main
import (
    "LearningGo/Functions_methods_interfaces/data"
)


func main(){
    data.PrintX()
    
    var p data.Point
    p.InitMe(2,4)
    p.PrintMe()
    p.Scale(3)
    p.PrintMe()
}
