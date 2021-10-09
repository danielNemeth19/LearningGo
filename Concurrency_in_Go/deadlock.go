package main
import (
    "sync"
)


var wg sync.WaitGroup


func dostuff(c1 chan int, c2 chan int) {
    <- c1
    c2 <- 1
    wg.Done()
}

/* Deadlock
 *  - deadlock is caused by circular dependencies
 *  - can be caused by waiting on channels (or mutex lock - unlock)
 *  - in this example each goroutine blocked on channel read
 *  - Golang runtime automatically detects when all goroutines are deadlocked
 *  - it cannot detect automatically when a subset of goroutines are deadlocked
 */
func main(){
    ch1 := make(chan int)
    ch2 := make(chan int)
    wg.Add(2)
    go dostuff(ch1, ch2)
    go dostuff(ch2, ch1)
    wg.Wait()
}
