package main

import "fmt"

func main() {
	a := 414
	b := 25

	fmt.Printf("before: a=%d b=%d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("after:  a=%d b=%d\n", a, b)
}
