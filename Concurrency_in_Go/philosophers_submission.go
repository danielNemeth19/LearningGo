package main
import (
    "fmt"
    "sync"
)


var wg sync.WaitGroup


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
        philos[i] = &Philo{leftCS:s[i], rightCS:s[(i+1)%5],num:i+1}
    }
    return philos
}


func (p Philo) eat(){
    for i:=0; i < 3; i++ {
        p.leftCS.Lock()
        p.rightCS.Lock()
        
        fmt.Printf("starting to eat %d (with chopstick %d and %d)\n", p.num, p.leftCS.serial_n, p.rightCS.serial_n)
        fmt.Printf("finishing eating %d (with chopstick %d and %d)\n", p.num, p.leftCS.serial_n, p.rightCS.serial_n)
        
        p.leftCS.Unlock()
        p.rightCS.Unlock()
        wg.Done()
    }
}


func main(){
    
    CSticks := make([] *Chops, 5)
    for i:=0 ; i < 5; i++ {
        chop := new(Chops)
        chop.serial_n = i
        CSticks[i] = chop
    }
    
    philos := create_philos_slice(CSticks)
    
    wg.Add(5)
    for i:=0; i < 5; i++ {
        go philos[i].eat()
    }
    wg.Wait()
}


