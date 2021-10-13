package main
import (
    "fmt"
//     "time"
)


func worker(queue chan int, worknumber int, done chan bool){
    for true {
        select {
            case k := <- queue:
                fmt.Println("doing work", k, "worknumber", worknumber)
                done <- true
        }
    }
}


func main(){
    q := make(chan int)
    done := make(chan bool)
    
    numberOfWorkers := 2
    for i := 0; i < numberOfWorkers; i++ {
        go worker(q, i, done)
    }
    
    numberOfJobs := 17
    for j := 0; j < numberOfJobs; j++ {
        go func(j int){
            q <- j 
        }(j)
    }
    
    for c := 0; c < numberOfJobs; c++ {
        <- done
    }
    fmt.Println("finished")
}
