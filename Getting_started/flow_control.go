package main
import "fmt"


func if_statement(){
    x:= 10
    if x > 5 {
        fmt.Printf("%v is greater than 5\n", x)
    }
    fmt.Printf("If statement func ended\n\n")
}


func loop_1(){
    for i:=0; i<10; i++ {
        fmt.Printf("Loop_1 -> Value of i: %v\n", i)
    }
    fmt.Printf("Loop_1 ended\n\n")
}


func loop_2(){
    k := 0
    for k < 10 {
        fmt.Printf("Loop_2 -> Value of k: %v\n", k)
        k++
    }
    fmt.Printf("Loop_2 ended\n\n")
}


func switch_case(){
    for y:=19; y >= 17; y-- {
        switch y {
            case 19:
                fmt.Printf("Case 19 executed as y is: %v\n", y)
            case 18:
                fmt.Printf("Case 18 executed as y is: %v\n", y)
            default:
                fmt.Printf("Default case executed as y is: %v\n", y)
        }
    }
    fmt.Printf("Switch case ended\n\n")
}


func tagless_switch(){
    for x:=0; x < 5; x++ {
        switch {
            case x < 2:
                fmt.Printf("Case 1 executes as x is: %v\n", x)
            case x < 3:
                fmt.Printf("Case 2 executes as x is: %v\n", x)
            default:
                fmt.Printf("Default executes as x is: %v\n", x)
        }
    }
    fmt.Printf("Tagless switch case ended\n\n")
}


func give_me_a_break(){
    for x :=0; x < 10; x++ {
        fmt.Printf("x is now: %v\n", x)
        if x == 5 {
            break 
        }
    }
    fmt.Printf("Break ended\n\n")
}


func continue_example(){
    y := 0
    for y < 5 {
        y++
        if y == 3 {
            continue
        }
        fmt.Printf("y is now: %v\n", y)
    }
    fmt.Printf("Continue ended\n\n")
}


func scan_user_input(){
    var myNumber int
    
    fmt.Printf("Enter a number:\n")
	
	num, err := fmt.Scan(&myNumber)
    if err != nil {
        fmt.Printf("%T, %v\n", err, err)
    }
    if err == nil {
        fmt.Printf("Got proper type for num: %T, %v\n", num, num)
    }
	
	fmt.Printf("Number entered: %v\n", myNumber)
    fmt.Printf("Scan ended\n\n")
}


func main(){
    if_statement()
    loop_1()
    loop_2()
    switch_case()
    tagless_switch()
    give_me_a_break()
    continue_example()
    scan_user_input()
}
