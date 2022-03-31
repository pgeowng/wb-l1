package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.RWMutex
	count int
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func (c *Counter) Get() int {
	c.RLock()
	defer c.RUnlock()
	return c.count
}

func main() {
	wg := sync.WaitGroup{}
	c := Counter{}

	for i := 0; i < 100; i++ {
		idx := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(idx%5))
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c.Get())
}
