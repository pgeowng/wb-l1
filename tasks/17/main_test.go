package main

import "testing"

func TestLowerBound(t *testing.T) {
	tests := []struct {
		arr      []int
		val      int
		expected int
	}{
		{[]int{}, -1, 0},
		{[]int{0}, -1, 0},
		{[]int{0}, 0, 0},
		{[]int{0}, 1, 1},
		{[]int{0}, 2, 1},
		{[]int{1, 1, 3, 3, 3, 4}, -2, 0},
		{[]int{1, 1, 3, 3, 3, 4}, -1, 0},
		{[]int{1, 1, 3, 3, 3, 4}, 0, 0},
		{[]int{1, 1, 3, 3, 3, 4}, 1, 0},
		{[]int{1, 1, 3, 3, 3, 4}, 2, 2},
		{[]int{1, 1, 3, 3, 3, 4}, 3, 2},
		{[]int{1, 1, 3, 3, 3, 4}, 4, 5},
		{[]int{1, 1, 3, 3, 3, 4}, 5, 6},
		{[]int{1, 1, 3, 3, 3, 4}, 6, 6},
	}

	for _, test := range tests {
		result := LowerBound(test.arr, test.val)
		if result != test.expected {
			t.Logf("for %#v, expected %d, got: %d", test, test.expected, result)
			t.Fail()
		}
	}
}

func TestUpperBound(t *testing.T) {
	tests := []struct {
		arr      []int
		val      int
		expected int
	}{
		{[]int{}, -1, 0},
		{[]int{0}, -1, 0},
		{[]int{0}, 0, 1},
		{[]int{0}, 1, 1},
		{[]int{0}, 2, 1},
		{[]int{1, 1, 3, 3, 3, 4}, -2, 0},
		{[]int{1, 1, 3, 3, 3, 4}, -1, 0},
		{[]int{1, 1, 3, 3, 3, 4}, 0, 0},
		{[]int{1, 1, 3, 3, 3, 4}, 1, 2},
		{[]int{1, 1, 3, 3, 3, 4}, 2, 2},
		{[]int{1, 1, 3, 3, 3, 4}, 3, 5},
		{[]int{1, 1, 3, 3, 3, 4}, 4, 6},
		{[]int{1, 1, 3, 3, 3, 4}, 5, 6},
		{[]int{1, 1, 3, 3, 3, 4}, 6, 6},
	}

	for _, test := range tests {
		result := UpperBound(test.arr, test.val)
		if result != test.expected {
			t.Logf("for %#v, expected %d, got: %d", test, test.expected, result)
			t.Fail()
		}
	}
}
func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr      []int
		val      int
		expected bool
	}{
		{[]int{}, -1, false},
		{[]int{0}, -1, false},
		{[]int{0}, 0, true},
		{[]int{0}, 1, false},
		{[]int{0}, 2, false},
		{[]int{1, 1, 3, 3, 3, 4}, -2, false},
		{[]int{1, 1, 3, 3, 3, 4}, -1, false},
		{[]int{1, 1, 3, 3, 3, 4}, 0, false},
		{[]int{1, 1, 3, 3, 3, 4}, 1, true},
		{[]int{1, 1, 3, 3, 3, 4}, 2, false},
		{[]int{1, 1, 3, 3, 3, 4}, 3, true},
		{[]int{1, 1, 3, 3, 3, 4}, 4, true},
		{[]int{1, 1, 3, 3, 3, 4}, 5, false},
		{[]int{1, 1, 3, 3, 3, 4}, 6, false},
	}

	for _, test := range tests {
		result := BinarySearch(test.arr, test.val)
		if result != test.expected {
			t.Logf("for %#v, expected %v, got: %v", test, test.expected, result)
			t.Fail()
		}
	}
}
