package main
import (
    "fmt"
    "sync"
)


func foo(wg *sync.WaitGroup){
    fmt.Printf("New Routine\n")
    wg.Done()
}


func main(){
    var wg sync.WaitGroup
    
    wg.Add(1)
    go foo(&wg)
    wg.Wait()
    
    fmt.Printf("Main Routine\n")
}
