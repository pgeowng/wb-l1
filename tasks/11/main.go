package main

import "fmt"

// Go спецификация не говорит конкретную сложность на добавление, удаление и получение в map.
// Предпологается, что операции работают за амортизированую O(1)
// https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics
// http://www.cs.cornell.edu/courses/cs3110/2011sp/Lectures/lec20-amortized/amortized.htm

// Пересечение множеств подразумевает проверку равенства между элементами,
// поэтому можно использовать map.
func intersect(a []float64, b []float64) []float64 {
	if len(a) == 0 || len(b) == 0 {
		return []float64{}
	}

	ccap := len(a) + len(b)
	mm := make(map[float64]int, ccap)

	for _, val := range a {
		mm[val]++
	}

	for _, val := range b {
		mm[val]++
	}

	result := make([]float64, 0, ccap)
	for val, count := range mm {
		if count > 1 {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	a := []float64{-3.4, -2.5, -1.9, 0.2, 1.3}
	b := []float64{-3.9, -2.5, -1.0, 1.3, 2.5}
	fmt.Println(intersect(a, b))

	a = []float64{1.2, 3.2, 2.2}
	b = []float64{3.2, 2.2, 1.2}
	fmt.Println(intersect(a, b))

}
