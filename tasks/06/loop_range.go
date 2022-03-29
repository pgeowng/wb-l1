package main

import (
	"fmt"
	"time"
)

// Выход, по закрытию входного канала.
func routine(ch <-chan int) {
	defer fmt.Println("stopped")
	fmt.Println("working")

	// Альтернативный код для того же результата
	// for val := range ch {
	// 	_ = val
	// }

	for {
		_, ok := <-ch
		if !ok {
			return
		}
	}
}

func main() {
	ch := make(chan int)
	go routine(ch)

	time.Sleep(time.Second)
	close(ch)
	time.Sleep(time.Second)
}
