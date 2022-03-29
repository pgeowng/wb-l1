package main

import (
	"fmt"
	"time"
)

// Выход если можем считать канал
func routine(done <-chan struct{}) {
	defer fmt.Println("stopped")
	fmt.Println("working")
	for {
		select {
		case <-done:
			return
		}
	}
}

func main() {

	done := make(chan struct{})
	go routine(done)

	time.Sleep(time.Second)
	close(done)
	time.Sleep(time.Second)
}
