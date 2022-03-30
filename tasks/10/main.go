package main

import "fmt"

// Математическое округление - округление в левую строну
//   math.Floor(29.8) -> 29
//   math.Floor(-29.8) -> -30
// Округление в сторону нуля
//   int(29.8) -> 29
//   int(-29.8) -> -29
func main() {
	input := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -29.8}

	mm := map[int][]float64{}

	for _, val := range input {
		key := int(val) / 10 * 10
		mm[key] = append(mm[key], val)
	}

	fmt.Printf("%v\n", mm)
}
