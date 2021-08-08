package main
import (
    "fmt"
    "math"
)

/* Functions can return functions
  - function with contollable parameters can be created
  - below example can create a function, that calculates distance from origin, based on a defined origin
  - origin location is passed as an argument - and built-in to the returned function
  - lexical scoping: environment includes names defined in block where the function is defined
  - closure: function + its environment -> when functions are passed/returned, their environment comes with them
  - o_y and o_y are in the closure of fn()
 */

func MakeDistOrigin (o_x, o_y float64) func (float64, float64) float64 {
    fn := func (x, y float64) float64 {
        return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
    }
    return fn
}


func main(){
    Dist1 := MakeDistOrigin(0, 0)
    fmt.Println(Dist1(2,2))
    
    Dist2 := MakeDistOrigin(2, 2)
    fmt.Println(Dist2(2,2))
}
