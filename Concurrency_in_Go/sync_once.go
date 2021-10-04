package main
import (
    "fmt"
    "sync"
)


var wg sync.WaitGroup
var on sync.Once


func setup(){
    fmt.Println("Initialization happens...")
}


func dostuff(){
    on.Do(setup)
    fmt.Println("dostuff is running")
    wg.Done()
}

/*Synchronous Initialization:
 *  - must happen once
 *  - must happen before everything else
 *  - use sync.Once -> once.Do(f) method:
 *      - function f will only be executed once, even if called in multiple go routines
 *      - all calls to once.Do() block until the first one returnes
 *      - ensures that initialization executes first and only one time
 */ 

func main(){
    wg.Add(2)
    go dostuff()
    go dostuff()
    wg.Wait()
}
