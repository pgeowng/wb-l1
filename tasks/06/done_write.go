package main

import (
	"fmt"
	"time"
)

// Выход если можем записать в канал без блокировки
func routine(done chan<- struct{}) {
	defer fmt.Println("stopped")
	fmt.Println("working")
	for {
		select {
		case done <- struct{}{}:
			return
		}
	}
}

func main() {
	done := make(chan struct{})
	go routine(done)

	time.Sleep(time.Second)
	<-done
	time.Sleep(time.Second)
}
