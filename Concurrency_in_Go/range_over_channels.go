package main
import (
    "fmt"
)


func main(){
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)
    

    
/*Iterating through a channel
 * one iteration each time a new value is received from channel
 * 'elem' is assigned to the read value in every iteration
 * iteration stops when sender calls close(queue)
 * using this range construct is the only case when close need to be called on a channel
 */
    for elem := range(queue) {
        fmt.Println(elem)
    }
}
