package main
import (
    "fmt"
)


type Shape2D interface {
    Area() float64
    Perimeter() float64
}


type Triangle struct {
    base float64
    height float64
    a float64
    c float64
}


type Rectangle struct {
    width float64
    length float64
}


func (t Triangle) Area() float64 {
    area := 0.5 * t.base * t.height
    return area 
}


func (t Triangle) Perimeter() float64 {
    perimeter := t.base + t.a + t.c
    return perimeter
}


func (r Rectangle) Area() float64 {
    area := r.width * r.length
    return area
}


func (r Rectangle) Perimeter() float64 {
    perimeter := 2 * (r.width + r.length)
    return perimeter
}

/*
 * Function parameter is an interface
 *  - use it when function needs to take multiple types as a parameter
 *  - all types must satisfy the interface though: interface methods must be those needed by the function
 */
func FitInYard(s Shape2D) bool {
    if (s.Area() < 100 && s.Perimeter() < 100) {
        return true
    }
    return false
}

/*
 * Empty interface
 *  - specifies no methods
 *  - all types satisfies the empty interface
 *  - use it to have a function accept any type of paramter
 */
func PrintMe(val interface{}) {
    fmt.Println(val)
}


func main() {
  myt := Triangle{base:3, height:4, a: 4, c:6}
  myr := Rectangle{width:20, length:15}
  
  fmt.Println(FitInYard(myt))
  fmt.Println(FitInYard(myr))
  
  PrintMe(myt)
  PrintMe("Any type can be printed by this function")
  
}
