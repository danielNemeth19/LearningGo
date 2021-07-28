package main
import "fmt"

func question_1(){
    x := []int {4, 8, 5}
    y := -1
    for _, elt := range x {
        if elt > y {
            y = elt
        }
    }
    fmt.Println(y)
}

func question_2() {
    x := [...]int {4, 8, 5}
    y := x[0:2]
    z := x[1:3]
    y[0] = 1
    z[1] = 3
    fmt.Println(x)
}


func question_3() {
    x := [...]int {1, 2, 3, 4, 5}
    y := x[0:2]
    z := x[1:4]
    fmt.Println(len(y), cap(y), len(z), cap(z))
}


func question_4() {
    x := map[string]int {"ian": 1, "harris": 2}
    for i, j := range x {
        if i == "harris" {
        fmt.Println(i, j)
        }
    }
}

type P struct {
    x string
    y int
}


func question_5() {
    b := P{"x", -1}
    a := [...]P{P{"a", 10}, P{"b", 2}, P{"c", 3}}
    for _, z := range a {
        if z.y > b.y {
            b = z
        }
    }
    fmt.Println(b.x)
}


func question_6() {
  s := make([]int, 0, 3)
  s = append(s, 100)
  fmt.Println(len(s), cap(s))
}


func main(){
    question_1()
    question_2()
    question_3()
    question_4()
    question_5()
    question_6()
}
