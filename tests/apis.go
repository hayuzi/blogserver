package main

import (
	"fmt"
	"sync"
	"time"
)

var cond *sync.Cond

func main() {
	cond = sync.NewCond(&sync.Mutex{})
	wg := do(job1, job2, job3, job4, job5)
	wg.Wait()
}

func job1() {
	time.Sleep(time.Second * 1)
	fmt.Println("job1")
	cond.Broadcast()
}
func job2() {
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
	fmt.Println("job2")
}
func job3() {
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
	fmt.Println("job3")
}
func job4() {
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
	fmt.Println("job4")
}
func job5() {
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
	fmt.Println("job5")
}

func do(fs ...func()) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	for _, fn := range fs {
		wg.Add(1)
		go func(f func()) {
			f()
			wg.Done()
		}(fn)
	}
	return wg
}
