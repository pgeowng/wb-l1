package main

import (
	"testing"
)

type Test struct {
	input    []int
	idx      int
	expected []int
}

var tests []Test = []Test{
	{nil, -1, nil},
	{nil, 0, nil},
	{nil, 41, nil},
	{[]int{1}, -1, []int{1}},
	{[]int{1}, 0, []int{}},
	{[]int{1}, 41, []int{1}},
	{[]int{1, 2}, -1, []int{1, 2}},
	{[]int{1, 2}, 0, []int{2}},
	{[]int{1, 2}, 1, []int{1}},
	{[]int{1, 2}, 41, []int{1, 2}},
	{[]int{1, 2, 3}, -1, []int{1, 2, 3}},
	{[]int{1, 2, 3}, 0, []int{2, 3}},
	{[]int{1, 2, 3}, 1, []int{1, 3}},
	{[]int{1, 2, 3}, 2, []int{1, 2}},
	{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
	{[]int{1, 2, 3}, 41, []int{1, 2, 3}},
}

func TestRemove(t *testing.T) {
	for _, test := range tests {
		result := remove(test.input, test.idx)
		for idx := range result {
			if result[idx] != test.expected[idx] {
				t.Logf("for %v and idx %d expected %v, got %v\n", test.input, test.idx, test.expected, result)
				t.Fail()
				continue
			}
		}
	}
}
