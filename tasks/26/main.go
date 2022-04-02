package main

import (
	"fmt"
	"unicode"
)

func HasOnlyUniq(str string) bool {
	m := make(map[rune]struct{})

	for _, rune := range str {
		t := unicode.ToLower(rune)
		_, ok := m[t]
		if ok {
			return false
		}
		m[t] = struct{}{}
	}

	return true
}

func main() {
	var str string

	str = "abcd"
	fmt.Printf("%s: %v\n", str, HasOnlyUniq(str))

	str = "abCdefAaf"
	fmt.Printf("%s: %v\n", str, HasOnlyUniq(str))

	str = "aabcd"
	fmt.Printf("%s: %v\n", str, HasOnlyUniq(str))

	str = "abCdef"
	fmt.Printf("%s: %v\n", str, HasOnlyUniq(str))

	str = "abCdefA"
	fmt.Printf("%s: %v\n", str, HasOnlyUniq(str))

	str = ""
	fmt.Printf("%s: %v\n", str, HasOnlyUniq(str))

	str = "ÑN"
	fmt.Printf("%s: %v\n", str, HasOnlyUniq(str))

	str = "\u0410\u0041"
	fmt.Printf("%s: %v\n", str, HasOnlyUniq(str))

	str = "\u767d\u9d6c\u7fd4"
	fmt.Printf("%s: %v\n", str, HasOnlyUniq(str))
}
