package main

import(
    "fmt"
    "sync"
)

type Host struct {
    count int
    MAX int
    permissionLock sync.Mutex
    countLock sync.Mutex
}

func (h Host)  getPermission(wg *sync.WaitGroup) {
    h.countLock.Lock()
    if h.count < h.MAX {
        h.count += 1
    }
    if h.count >=  h.MAX {
       // fmt.Println("Locking permission at count ", h.count)
        h.permissionLock.Lock()
    }
   // fmt.Println("Grantting permission at count ", h.count)
    h.countLock.Unlock()
    wg.Done()
}

func (h Host) releasePermission() {
    h.countLock.Lock()
    if (h.count >= h.MAX) {
        h.count -= 1
//        fmt.Println("Unlocking permissions at count ", h.count)
        h.permissionLock.Unlock()
    }
 //   fmt.Println("Releasing permission at count", h.count)
    h.countLock.Unlock()
}

var host *Host = &Host{MAX: 2}

type ChopStick struct{
    sync.Mutex
}

type Philosopher struct{
    right, left *ChopStick
    id, count int
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
    var host_wg sync.WaitGroup
    for {
        if (p.count < 3) {
            host_wg.Add(1)
            go host.getPermission(&host_wg)
            host_wg.Wait()
            p.right.Lock()
            p.left.Lock()
            fmt.Println("Starting to eat ", p.id)
            fmt.Println("Finishing eating ", p.id)
            p.count += 1
            p.right.Unlock()
            p.left.Unlock()
            go host.releasePermission()
        } else {
            break
        }
    }
    wg.Done()
}


func main() {
    size := 5
    chops := make([]*ChopStick, size)
    for i := 0; i < size; i++ {
        chops[i] = new(ChopStick)
    }
    phil := make([]*Philosopher, size)
    for i := 0; i < size; i++ {
        phil[i] = &Philosopher{right: chops[i], left: chops[(i+1) % size], id: i+1, count: 0}
    }
    var wg sync.WaitGroup
    for _, p := range phil {
        wg.Add(1)
        go p.eat(&wg)
    }
    wg.Wait()
}
