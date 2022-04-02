package main

import "fmt"

func main() {
	arr := []int{0, 1, 1, 2, 3, 5, 8, 11}

	// Используем copy, если удаление из середины.
	idx := 5
	copy(arr[idx:], arr[idx+1:])
	// Всегда сокращаем длину
	arr = arr[:len(arr)-1]

	fmt.Println(arr)

	// Eсли порядок неважен, то можно обойтись без copy.
	idx = 2
	arr[idx] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]

	fmt.Println(arr)
}
