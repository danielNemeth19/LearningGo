package main
import (
    "fmt"
    "math"
    "os"
)


func GenDisplaceFn(acc, v_init, s_init float64) func(time float64) float64 {
    fn := func (time float64) float64 {
        return 0.5*acc*math.Pow(time, 2) + v_init*time + s_init
    }
    return fn
}


func get_param(param string) float64 {
    var param_value float64
    fmt.Printf("Specify value for %s: ", param)
    _, err := fmt.Scanf("%f", &param_value)
    if err != nil {
        fmt.Printf("Conversion error: %v\n", err)
        os.Exit(1)
    }
    return param_value
}


func main(){
    acc := get_param("acceleration")
    v_init := get_param("initial velocity")
    s_init := get_param("initial displacement")
    
    fn := GenDisplaceFn(acc, v_init, s_init)
    t := get_param("time")
    res := fn(t)
    fmt.Printf("Displacement after elapsed time %v is: %v\n", t, res)
}
