package main
import (
    "fmt"
)

/*
 * Interfaces
 * Set of method signatures
 *  - name, parameters, return values
 *  - implementation is NOT defined
 * Used to express conceptual similarity between types
 * Example: Shape2D interface - all 2D shapes must have Area() and Perimeter()
 */
type Shape2D interface {
    Area() float64
    Perimeter() float64
}


/*
 * Satisfying an Interface
 * Type satisfies an interface if type defines all methods specified in the interface
 *  - same method signatures needed
 *  - additional methods (and data) are OK
 * Having Interface in go similar to inheritance with overriding
 */
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


func main() {
    myt := Triangle{base:3, height:4, a: 4, c:6}
    fmt.Println(myt.Area())
    fmt.Println(myt.Perimeter())

// this technique can be used to check if Triangle satisfies the interface Shape2D    
    var i interface{} = myt
    _, flag := i.(Shape2D)
    fmt.Printf("Triangle satisfies the Shape2D interface: %v\n", flag)
    
    myr := Rectangle{width:2, length:4}
    fmt.Println(myr.Area())
    fmt.Println(myr.Perimeter())
    
    var i2 interface{} = myr
    _, flag2 := i2.(Shape2D)
    fmt.Printf("Rectangle satisfies the Shape2D interface: %v\n", flag2)
}



