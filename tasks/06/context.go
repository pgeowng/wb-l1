package main

import (
	"context"
	"fmt"
	"time"
)

// Выход, если контекст завершен.
func routine(ctx context.Context) {
	defer fmt.Println("stopped")
	fmt.Println("working")

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go routine(ctx)

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}
