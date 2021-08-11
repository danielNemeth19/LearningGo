package main
import (
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
    "bufio"
)


func GenDisplaceFn(acc, v_init, s_init float64) func(time float64) float64 {
    fn := func (time float64) float64 {
        return 0.5*acc*math.Pow(time, 2) + v_init*time + s_init
    }
    return fn
}


func get_func_parameters() (float64, float64, float64) {
    fmt.Printf("Define values for acceleration, initial velocity and initial displacement (separated by space): ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    inputString := scanner.Text()
    inputNumbers := strings.Split(inputString, " ")
        
    if len(inputNumbers) != 3 {
        fmt.Printf("Three numbers need to be specified: got %d\n", len(inputNumbers))
        os.Exit(1)
    }
    
    acc := _parse_as_float(inputNumbers[0])
    v_init := _parse_as_float(inputNumbers[1])
    s_init := _parse_as_float(inputNumbers[2])
    return acc, v_init, s_init
}


func get_time_param() float64 {
    var t float64
    fmt.Printf("Specify the time: ")
    _, err := fmt.Scanf("%f", &t)
    if err != nil {
        fmt.Printf("Conversion error: %v\n", err)
        os.Exit(1)
    }  
    return t
}


func _parse_as_float(s string) float64{
    number, err := strconv.ParseFloat(s, 64)
    if err != nil {
        fmt.Printf("Conversion error: %v\n", err)
        os.Exit(1)
    }
    return number
}

func main(){
    acc, v_init, s_init := get_func_parameters()
    fn := GenDisplaceFn(acc, v_init, s_init)
    t := get_time_param()
    res := fn(t)
    fmt.Printf("Displacement after elapsed time %v is: %v\n", t, res)
}
