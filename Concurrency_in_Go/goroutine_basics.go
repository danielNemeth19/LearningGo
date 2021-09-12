package main

import (
    "fmt"
    "time"
)


/*
 * Only "Main routine" will be printed
 * this is because main finished before the new goroutine started
 * A goroutine exits when its code is complete:
 *  - when the main goroutine is complete (which is created automaically) all other goroutines exits
 * Creating a goroutine means, the main function is not blocked on the given call
 */

func main(){
    go fmt.Printf("New routine\n")
    
    /*
     * A hackish way to solve the problem: delayed exit
     * Waiting in the main routine will cause the run time scheduler to schedule the other goroutine
     * However this is only an assumption! Timing is nondeterministic:
     *      - maybe the OS schedules (which runs above the Go Runtime scheduler) an another thread
     *      - maybe the Go Runtime scheduler schedules another goroutine
     */
    time.Sleep(100 * time.Millisecond)
    
    fmt.Printf("Main routine\n")
}
