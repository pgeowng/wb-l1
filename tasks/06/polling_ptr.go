package main

import (
	"fmt"
	"time"
)

// Выход, если условие сработало
func routine(cond *bool) {
	defer fmt.Println("stopped")
	fmt.Println("working")

	// постоянно вызывает функцию cond
	// решением может быть интервал между запросами
	for {
		if *cond {
			return
		}
	}
}

func main() {
	stop := false
	go routine(&stop)

	time.Sleep(time.Second)
	stop = true
	time.Sleep(time.Second)
}
