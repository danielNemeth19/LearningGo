package main

import (
	"fmt"
    "time"
)


func add2(i int) {
    fmt.Printf("result is %d\n", i+2)
}

func sub2(i int) {
    fmt.Printf("result is %d\n", i-2) 
}


/*
 * Race condition occurs because:
 *  - there's one variable seen by both goroutines (num)
 *  - both goroutines prints a modified version of the given number:
 *      - one of them print x + 2
 *      - the other prints  x - 2
 *  - however since the order of executions are not deterministic this program cannot guarantee the order of the call functions
 
 */

func main(){
    num := 5
    go add2(num)
    go sub2(num)
    time.Sleep(1 * time.Second)
}
