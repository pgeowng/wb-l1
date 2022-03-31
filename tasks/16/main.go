package main

import (
	"fmt"
	"math/rand"
	"time"
)

var showSwap bool = true

// Быстрая сортировка Хоара
// Основной смысл заключается в том, чтобы разделить массив на 2 части.
// Одна больше x, другая меньше x и отсортировать эти части независимо друг от друга.
// А затем объединить их вместе.

// [0 5 3 2]
// [0 2 3 5]
// [0] и [3 5]

func sort(arr []int, left int, right int) {
	// При попытке отсортировать [] или [ 1 ], сразу говорим, что они отсортированы.
	if right-left <= 0 {
		return
	}

	// Выбираем случайный элемент, чтобы минимизировать предсказуемость сортировки.
	// Если использовать конкретный способ выбора pivot элемента, то можно с легкостью сгенерировать худший случай сортировки
	rnd := rand.Intn(right-left+1) + left
	arr[rnd], arr[right] = arr[right], arr[rnd]

	pivot := arr[right]
	mid := left

	// Пробигаем по массиву. Слево от mid оставляем элементы, которые меньше pivot
	for i := left; i < right; i++ {
		if arr[i] < pivot {
			arr[mid], arr[i] = arr[i], arr[mid]
			mid++
		}
	}

	// Мы знаем, что слево от mid элементы меньше pivot.
	// Поэтому pivot может занять свою конечную позицию.
	arr[mid], arr[right] = arr[right], arr[mid]
	sort(arr, left, mid-1)
	sort(arr, mid+1, right)

}

func QuickSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	sort(input, 0, len(input)-1)
	return input
}

func main() {
	rand.Seed(time.Now().UnixNano())
	input := []int{4, 2, 6, 23, 1, 5, 2}

	fmt.Println(QuickSort(input))
}
