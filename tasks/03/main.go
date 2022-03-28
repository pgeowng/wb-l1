package main

import "fmt"

func main() {
	input := []int{2, 4, 6, 8, 10}

	// Закрываем канал, чтобы обозначить окончание данных.
	// Тем самым говорим, что окончательный результат может быть получен.
	in := make(chan int, len(input))
	go func() {
		defer close(in)
		for _, num := range input {
			in <- num
		}
	}()

	out := make(chan int)
	go func() {
		defer close(out)
		result := 0
		for n := range in {
			result += n * n
		}
		out <- result
	}()

	fmt.Println(<-out)
}
