package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	number          int
}

func (p Philo) eat(c chan int) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		//acquireing the semaphore (channel)
		c <- 1
		if rand.Intn(2) == 0 {
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Printf("starting to eat %d\n", p.number)
			time.Sleep(100 * time.Millisecond)
			p.rightCS.Unlock()
			p.leftCS.Unlock()
			fmt.Printf("finishing eating %d\n", p.number)
			//else pick right first
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
			fmt.Printf("starting to eat %d\n", p.number)
			time.Sleep(100 * time.Millisecond)
			p.leftCS.Unlock()
			p.rightCS.Unlock()
			fmt.Printf("finishing eating %d\n", p.number)
		}
		//releasing the channel
		<-c
	}
}

//main is the HOST by defining the channel with buffer size = 2
func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}
	wg.Add(5)

	//a buffered channel of size two for allowing no more than two concurrent event
	c := make(chan int, 2)
	for i := 0; i < 5; i++ {
		go philos[i].eat(c)
	}
	wg.Wait()

}
