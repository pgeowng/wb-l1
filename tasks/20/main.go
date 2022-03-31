package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func IsSpaceStr(str string) bool {
	for _, r := range str {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// Обычный реверс: "\thello  world " -> " world  \thello"
// Реверс слов:    "\thello  world " -> "\tworld  hello "
func ReverseWords(str string) string {
	tokens := []string{}
	start := 0
	newWord := false

	// Рассматриваем каждый символ.
	// Нулевой игнорируем.
	// Если встретили пробел, то начинаем новый токен.
	for idx, r := range str {
		switch {
		case idx == 0:
			newWord = unicode.IsSpace(r)
		case unicode.IsSpace(r):
			tokens = append(tokens, str[start:idx])
			start = idx
			newWord = true
		case newWord:
			tokens = append(tokens, str[start:idx])
			start = idx
			newWord = false
		}
	}
	tokens = append(tokens, str[start:])

	// Пробегаем с обоих концов, до встречи указателей. Если нужно свапим.
	size := len(tokens)
	left := 0
	right := size - 1
	for left < right {
		for left < right && IsSpaceStr(tokens[left]) {
			left++
		}

		for left < right && IsSpaceStr(tokens[right]) {
			right--
		}

		if left < right {
			tokens[left], tokens[right] = tokens[right], tokens[left]
			left++
			right--
		}
	}

	return strings.Join(tokens, "")

}

// echo "snow dog sun" | go run main.go
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(ReverseWords(scanner.Text()))
	}
}
