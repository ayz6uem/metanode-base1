package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Count() {
	mutexCounter()
	atomicCounter()
}

func mutexCounter() {
	c := Counter{}
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				c.Inc()
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("task2 counter count", c.i)
}

type Counter struct {
	i      int
	locker sync.Mutex
}

func (c *Counter) Inc() {
	c.locker.Lock()
	c.i++
	c.locker.Unlock()
}

func atomicCounter() {
	ai := atomic.Int64{}

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				ai.Add(1)
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("task2 counter atomic count", ai.Load())
}
