package main

import "fmt"

// Функции LowerBound и UpperBound позволяют найти тот диапозон,
// в котором находятся несколько значений value.
// Происходит рассмотрение значений в диапозоне (-1 .. len(arr)] // () - невключено, [] - включено
// Результат (left+right)/2 находится всегда [0..len(arr)-1]
// Такой поиск удобен если нужно пробежаться по нескольким элементам с определенным ключем.
//   left := LowerBound(arr, 2)
//   right := UpperBound(arr, 2)
//   for i := left; i < right; i++ {...}

// Возвращает левый индекс, значение которого равно или больше value.
// Если все значения < value, тогда возвращает len(arr)
func LowerBound(arr []int, value int) (right int) {
	left := -1
	right = len(arr)

	for right-left > 1 {
		mid := (left + right) / 2
		if arr[mid] < value {
			left = mid
		} else {
			right = mid
		}
	}

	return
}

// Возвращает левый индекс, значение которого больше value.
// Если все значения <= value, тогда вернет len(arr)
func UpperBound(arr []int, value int) (right int) {
	left := -1
	right = len(arr)
	for right-left > 1 {
		mid := (left + right) / 2
		if arr[mid] <= value {
			left = mid
		} else {
			right = mid
		}
	}

	return
}

func BinarySearch(arr []int, value int) bool {
	left := -1
	right := len(arr)

	for right-left > 1 {
		mid := (left + right) / 2
		if arr[mid] == value {
			return true
		} else if arr[mid] < value {
			left = mid
		} else {
			right = mid
		}
	}

	return false
}

func main() {
	input := []int{0, 2, 2, 2, 4, 4, 5, 5, 5, 8, 10, 10}

	left := LowerBound(input, 2)
	right := UpperBound(input, 2)
	fmt.Println(input[left:right], left, right)

	val := -1
	found := BinarySearch(input, val)
	fmt.Println(val, "found?", found)

}
