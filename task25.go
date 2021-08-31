package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.RWMutex
	count int
	// count int32
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func (c *Counter) Get() int {
	c.RLock()
	defer c.RUnlock()
	return c.count
}

// func (c *Counter) Get() int {
// 	for {
// 		val := atomic.LoadInt32((&c.count))
// 		if atomic.CompareAndSwapInt32(&c.count, val, val+1) {
// 			return val
// 		}
// 	}
// }

// func (c *Counter) inc() {
// 	atomic.AddInt32(&c.count, 1)
// }

// func (c *Counter) get() int32 {
// 	return atomic.LoadInt32(&c.count)
// }

func main() {
	test(100)
}

func test(n int) {
	var c Counter
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			c.Increment()
			fmt.Println(c.Get()) // increment и get могут выдавать разные значения
			wg.Done()
		}()
	}
	wg.Wait()
}
