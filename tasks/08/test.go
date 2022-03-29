package main

import "fmt"

func SetBit(number int64, bit int, value int) int64 {
	number = number & ^(int64(1)<<bit) | (int64(value) << bit)
	return number
}

func main() {
	cases := []struct {
		num      int64
		bit      int
		val      int
		expected int64
	}{
		{0, 0, 0, 0},
		{0, 0, 1, 1},
		{1, 0, 1, 1},
		{1, 0, 0, 0},
	}

	for _, c := range cases {
		res := SetBit(c.num, c.bit, c.val)
		if res != c.expected {
			fmt.Printf("for %v expected %d, got: %d\n", c, c.expected, res)
		}
	}
}
