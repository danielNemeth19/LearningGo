package main
import (
    "fmt"
    "sync"
)


var i int = 0
var wg sync.WaitGroup
var mut sync.Mutex



/* Mutex (Mutual exclusion)
 *  - restricts possible interleavings
 *  - access to shared variables cannot be interleaved
 *  - mutex uses a binary semaphore (shared variable either in use or not)
 *  - methods Lock and Unlock handles this -> when Unlock is called, a blocked Lock() can proceed
 */ 
func inc(){
    mut.Lock()
    i = i + 1
    mut.Unlock()
    wg.Done()
}


func main(){
    increments := 3000
    wg.Add(increments)
     
    for x:=0; x <increments; x++ {
        go inc()
    }
    wg.Wait()
    
    fmt.Printf("This Program IS concurrency-safe!\n")
    fmt.Printf("Required result is: %d\n", increments)
    if i == increments {
        fmt.Printf("Succes by chance as i is: %d\n", i)
    } else {
        fmt.Printf("Failure: i is %d -- instead of %d\n", i, increments)
    }
}

