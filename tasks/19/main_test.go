package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"ababa", "ababa"},
		{"ababab", "bababa"},
	}

	for _, test := range tests {
		result := Reverse([]rune(test.input))
		if string(result) != test.expected {
			t.Logf("for %v expected %v, got: %v", test, test.expected, string(result))
			t.Fail()
		}
	}
}
