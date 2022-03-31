package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Reverse(arr []rune) []rune {
	last := len(arr) - 1
	mid := len(arr) / 2
	for i := 0; i < mid; i++ {
		arr[i], arr[last-i] = arr[last-i], arr[i]
	}
	return arr
}

// Считывает построчно stdin. Переворачивает каждую строку.
// Проверка:
//   $ cat main.go | go run main.go
//   $ echo "hello\nworld\nглаврыба\n日本語" | go run main.go
func main() {
	reader := bufio.NewReader(os.Stdin)

	line := make([]rune, 0, 10)

	for {
		r, _, err := reader.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "reverse: %v\n", err)
			os.Exit(1)
		}

		if r == '\n' {
			fmt.Println(string(Reverse(line)))
			line = line[:0]
		} else {
			line = append(line, r)
		}

	}

	fmt.Print(string(Reverse(line)))
}
