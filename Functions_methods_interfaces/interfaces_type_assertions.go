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


func SpecialTriangleMethod(t Triangle) {
    fmt.Printf("SpecialTriangleMethod has been called >> Triangle: %v\n", t)
}


func SpecialRectangleMethod(r Rectangle) {
    fmt.Printf("SpecialRectangleMethod has been called >> Rectangle: %v\n", r)
}

/*
 * Type assertion can be used to determine and extract the underlying concrete type
 * concrete type in parentheses
 * if interface contains concrete type:
 *  - rect == concrete type, ok == true
 * if interface does not contain concrete type:
 * - rect == zero, ok == false
 */
func SpecialShapeMethod(s Shape2D) bool {
    tri, ok := s.(Triangle)
    if ok {
        SpecialTriangleMethod(tri)
    }
    rect, ok := s.(Rectangle)
    if ok {
        SpecialRectangleMethod(rect)
    }
    return ok
}

/*
 * Switch statement used with a type assertion
 */
func SepcialMethodWithSwitch(s Shape2D) bool {
    switch sh := s.(type) {
        case Triangle:
            SpecialTriangleMethod(sh)
        case Rectangle:
            SpecialRectangleMethod(sh)
    }
    return true
}


func main(){
    myt := Triangle{base:3, height:4, a: 4, c:6}
    myr := Rectangle{width:20, length:15}
    
    SpecialShapeMethod(myt)
    SpecialShapeMethod(myr)
    
    SepcialMethodWithSwitch(myt)
    SepcialMethodWithSwitch(myr)
}

