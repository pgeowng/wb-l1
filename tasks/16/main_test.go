package main

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{},
		{[]int{0}, []int{0}},
		{[]int{0, 1}, []int{0, 1}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{1, 2, 0}, []int{0, 1, 2}},
		{[]int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}},
		{[]int{2, 2, 2, 2, 1, 1, 1}, []int{1, 1, 1, 2, 2, 2, 2}},
		{[]int{4, 2, 6, 23, 1, 5, 2}, []int{1, 2, 2, 4, 5, 6, 23}},
	}

	for _, test := range tests {
		result := QuickSort(test.input)
		if len(result) != len(test.output) {
			t.Logf("expected %v, got bad length: %v", test, result)
			t.Fail()
			continue
		}
		for idx := range result {
			if test.output[idx] != result[idx] {
				t.Logf("expected %v, got bad value for idx: %d. result %d != %d", test, idx, result[idx], test.output[idx])
				t.Fail()
				break
			}
		}
	}
}
