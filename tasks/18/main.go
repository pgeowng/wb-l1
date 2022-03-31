package main

import (
	"fmt"
	"math/rand"
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
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	c := Counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))
			c.Inc()
		}()
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case <-done:
			fmt.Println(c.Get())
			return
		case <-time.After(time.Duration(300) * time.Millisecond):
			fmt.Println(c.Get())
		}
	}
}
