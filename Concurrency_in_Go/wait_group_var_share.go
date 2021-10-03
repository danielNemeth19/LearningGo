package main
import (
    "fmt"
    "sync"
)


var i int = 0
var wg sync.WaitGroup

func inc(){
    i = i + 1
    wg.Done()
}


func main(){
    increments := 3000
    wg.Add(increments)
    
/*Goroutines sharing variables:
 *  Note - concurrency is on the machine code level! 
 *  This means that the interleavings of instructions happen on machine code level - not source code level
 *  Increment is at least 3 machine code instructions: 
 *      - it's possible that two goroutines read the same value as i and writes back the same value
 *      - i.e. some increment will be skipped..
 */ 
    
    for x:=0; x <increments; x++ {
        go inc()
    }
    wg.Wait()
    
    fmt.Printf("Program is not concurrency-safe, but could work by chance..\n")
    fmt.Printf("Required result is: %d\n", increments)
    if i == increments {
        fmt.Printf("Succes by chance as i is: %d\n", i)
    } else {
        fmt.Printf("Failure: i is %d -- instead of %d\n", i, increments)
    }
}

