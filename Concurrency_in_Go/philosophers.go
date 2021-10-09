package main
import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

type Chops struct{
    sync.Mutex
}


type Philo struct {
    leftCS, rightCS *Chops
}

/*Deadlock problem:
 *  - all philosophers might lock their left chopsticks concurrently
 *  - this would mean all chopsticks are locked
 *  - so locking the first right chopstick will not be possible!!
 */
func (p Philo) eat(n int){
    for {
        p.leftCS.Lock()
        p.rightCS.Lock()
        
        fmt.Printf("Philo %d is eating...\n", n)
        
        p.leftCS.Unlock()
        p.rightCS.Unlock()
        wg.Done()
    }
}


func main(){
    CSticks := make([] *Chops, 5)
    for i:=0 ; i < 5; i++ {
        CSticks[i] = new(Chops)
    }
    
    philos := make([] *Philo, 5)
    for i:=0 ; i < 5; i++ {
        philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
        fmt.Printf("Philosopher %d, picks up left: %d right %d\n", i, i, (i+1)%5)
    }
    
    wg.Add(5)
    for i:=0; i < 5; i++ {
        go philos[i].eat(i)
    }
    wg.Wait()
}
