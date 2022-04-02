package main

import (
	"fmt"
	"time"
)

func Sleep(dur time.Duration) {
	ns := dur.Nanoseconds()
	start := time.Now()
	for time.Now().Sub(start).Nanoseconds()-ns < 0 {
	}
}

func main() {
	timeout := 1

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("hello", i)
			Sleep(time.Duration(250) * time.Millisecond)
		}
	}()

	start := time.Now()
	Sleep(time.Duration(timeout) * time.Second)
	end := time.Now()
	fmt.Println("Sleep for:", timeout)
	fmt.Println("Time:", end.Sub(start))
}
