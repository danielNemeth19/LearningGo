package main
import (
    "fmt"
)

// variable number of args defined with ... -> they will be treated as a slice inside the funtion
func getMax(vals ...int) int {
    maxV := -1
    for _, v := range vals {
        if v > maxV {
            maxV = v
        }
    }
    return maxV
}


func main(){
    gn := getMax(23, 54, 6, 7, 1, 5, 6, 7)
    fmt.Println(gn)
    
    // slice can be passed in too, but needs ... as suffix
    vslice := []int{1 ,2, 10, 3, 5}
    fmt.Println(getMax(vslice...))
}
