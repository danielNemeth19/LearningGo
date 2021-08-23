package main
import (
    "fmt"
)


/*
* Interface type:
* Specified method signatures
* Implementations are abstracted
*/
type Speaker interface {
    Speak()
}


/*
* Concrete type:
* Specify the exact representation of the data and methods
* Complete method implementation included
*/
type Dog struct {
    name string
}


func (d *Dog) Speak(){
    if d == nil {
        fmt.Printf("No dog name has been defined: dynamic value is nil\n")
    } else {
        fmt.Printf("Name of my dog is: %s\n", d.name)
    }
}


/*
* Interface values:
* Can be treated like any other value (assigned, passed, returned)
* Interface values has two components:
*   - dynamic type: concrete type which is assigned to
*   - dynamic value: value of the dynamic type
*   - in this example dynamic type is Dog, dynamic value is myd
*/
func make_interface_value(){
    var s Speaker
    myd := &Dog{name:"Brian"}
    s = myd
    s.Speak()
}

/*
* Interface can have nil dynamic value
* Below 's' has a dynamic type, but no value
* This is enough to be able to call the method as type known -> implementation is given
*/
func make_interface_nil_dynamic_value(){
    var s Speaker
    var d *Dog
    s = d
    s.Speak()
}


/*
* Interface with nil dynamic type
*  - this is called: nil interface value
*  - very different from an interface with nil dynamic value
*  - cannot call method as type is not known -> hence no implentation
*/
func make_interface_nil_dynamic_type(){
    var s Speaker
    fmt.Printf("s.Speak() cannot be called here (%v)\n", s)
}


func main(){
    make_interface_value()
    make_interface_nil_dynamic_value()
    make_interface_nil_dynamic_type()
}







