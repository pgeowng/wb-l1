package main

import "fmt"

// Функция специально возвращает slice,
// чтобы показать, что передаваемый slice уже не валиден.
// Поведение схоже с append.
// Если удаляем последный элемент, то просто вернем slice с изменённым размером.
func remove(arr []int, idx int) []int {
	size := len(arr)
	if size == 0 || idx < 0 || idx >= size {
		return arr
	}
	nsize := size - 1
	if idx != nsize {
		copy(arr[idx:nsize], arr[idx+1:])
	}
	return arr[:nsize]
}

func main() {
	arr := []int{0, 1, 1, 2, 3, 5, 8, 11}
	fmt.Println("arr: ", arr)

	brr := remove(arr, 2)
	fmt.Println("arr: ", arr)
	fmt.Println("brr: ", brr)

	brr = remove(brr, 6)
	fmt.Println("brr: ", brr)
}
