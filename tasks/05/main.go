package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func readerSilent(ch <-chan int) {
	for {
		_ = <-ch
	}
}

func reader(ch <-chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func writer(ch chan<- int) {
	for {
		ch <- rand.Intn(1000000000)
	}
}

func main() {
	timeout := flag.Int("t", 5, "timeout in seconds")
	silent := flag.Bool("s", false, "silent")
	flag.Parse()
	if *timeout <= 0 {
		fmt.Println("timeout must be positive integer:", *timeout)
		return
	}

	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	go writer(ch)

	if *silent {
		go readerSilent(ch)
	} else {
		go reader(ch)
	}

	// Блокируем main перед выходом на N секунд
	// Просто для теста замеряем сколько было ожидание
	start := time.Now()
	time.Sleep(time.Duration(*timeout) * time.Second)
	end := time.Now()
	fmt.Println("Sleep for:", *timeout)
	fmt.Println("Time:", end.Sub(start))
}
