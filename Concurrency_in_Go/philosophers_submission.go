package main
import (
    "fmt"
    "sync"
)


type Chops struct{
    sync.Mutex
    serial_n int
}


type Philo struct {
    leftCS, rightCS *Chops
    num int
}


func create_philos_slice(s []*Chops) []*Philo {
    philos := make([] *Philo, 5)
    for i:= 0; i < 5; i++ {
        if i == 4 {
            philos[i] = &Philo{leftCS:s[0], rightCS:s[4],num:i+1}
        } else {
            philos[i] = &Philo{leftCS:s[i], rightCS:s[(i+1)%5],num:i+1}
        }
    }
    return philos
}


func (p Philo) eat(s chan bool, j int){
    for i:=0; i < 3; i++ {
        p.leftCS.Lock()
        p.rightCS.Lock()
        
        fmt.Printf("starting to eat %d (with chopstick %d and %d job %d)\n", p.num, p.leftCS.serial_n, p.rightCS.serial_n, j)
        fmt.Printf("finishing eating %d (with chopstick %d and %d job %d)\n", p.num, p.leftCS.serial_n, p.rightCS.serial_n, j)
        
        s <- true
        
        p.leftCS.Unlock()
        p.rightCS.Unlock()
    }
}


func host(signal chan bool, worknumber int, done chan bool){
    for true {
        select {
            case <- signal:
                done <- true
        }
    }
}


func main(){
    
    signal := make(chan bool)
    done := make(chan bool)
    
    CSticks := make([] *Chops, 5)
    for i:=0 ; i < 5; i++ {
        chop := new(Chops)
        chop.serial_n = i
        CSticks[i] = chop
    }
    
    philos := create_philos_slice(CSticks)
    
    numberOfWorkers := 2
    for i := 0; i < numberOfWorkers; i++ {
        go host(signal, i, done)
    }
    
    
    numberOfJobs := 15
    for j := 0; j < numberOfJobs; j++ {
        go philos[j%5].eat(signal, j)
    }
        
    for c := 0; c < numberOfJobs; c++ {
        <- done
    }
     
}


