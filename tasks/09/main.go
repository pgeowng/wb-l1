package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(30) + 10
	fmt.Printf("Sending %d values...\n", size)

	input := make([]int, 0, size)
	for i := 0; i < size; i++ {
		input = append(input, rand.Intn(100))
	}
	fmt.Println(input)

	inputCh := make(chan int, 5)
	go func() {
		defer close(inputCh)
		for _, val := range input {
			inputCh <- val
		}
	}()

	resultCh := make(chan int, 3)
	go func() {
		defer close(resultCh)
		for val := range inputCh {
			resultCh <- val * 2
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		idx := 1
		for val := range resultCh {
			fmt.Printf("%d: %d\n", idx, val)
			idx++
		}
	}()

	wg.Wait()
}
