package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello", "hello"},
		{"hello world", "world hello"},
		{"hello  world", "world  hello"},
		{"	hello  world", "	world  hello"},
		{"	hello  world	there", "	there  world	hello"},
		{"	another hello  world	there  	", "	there world  hello	another  	"},
	}

	for _, test := range tests {
		result := ReverseWords(test.input)
		if result != test.expected {
			t.Logf("for %#v expected %#v, got: %#v", test.input, test.expected, result)
			t.Fail()
		}
	}
}
