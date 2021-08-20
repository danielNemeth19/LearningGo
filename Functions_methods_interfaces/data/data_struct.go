package data
import (
    "fmt"
)


type Point struct {
    x float64
    y float64
}


/*
 * Pointer Receivers:
 * Receiver can be a pointer type
 * This will be a call by reference, as a pointer will be passed to the method
 *  - this mean the actual data in the receiver can be modified
 * Notice that within method:
 *  - there's not need to dereference object
 *  - i.e. Point object can be referenced as p and not as *p
 *  - syntactic sugar: dereferencing is automatic with . operator
 */
func (p *Point) InitMe(xn, yn float64) {
    p.x = xn
    p.y = yn
}


func (p *Point) Scale(v float64) {
    p.x = p.x * v
    p.y = p.y * v
}


func (p *Point) PrintMe() {
    fmt.Printf("x is: %v -- y is: %v\n", p.x, p.y)
}


func (p *Point) OffsetX(offx float64) {
    p.x = p.x + offx
}



/*
 * This is a call by value type of definition
 * Receiver is passed implicitly as an argument
 * The method won't be able to modify the actual data inside the receiver..
 * ..as it's only operating on a copy of the object
 * Other problem: if the receiver is large, copying the whole object can be problematic
 */
func (p Point) OffsetXBad(offx float64) {
    p.x = p.x + offx
}
