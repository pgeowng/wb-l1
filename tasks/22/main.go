package main

import (
	"errors"
	"fmt"
)

type Digit int

const base = 1000000000

// func add(a, b []Digit) []Digit {

// }

// func toDigits(n string) []Digit {

// }

func calc(op string, a string, b string) (result *string, err error) {
	switch op {
	case "+":
	case "-":
	case "*":
	case "/":
	default:
		return nil, errors.New("unknown operation")
	}
	return nil, nil
}

func validateNum(n string) error {
	for _, rune := range n {
		if rune < '0' || '9' < rune {
			return errors.New("number contains non digit rune: " + string(rune))
		}
	}
	return nil
}

func main() {
	var op, a, b string

	_, err := fmt.Scanf("%s %s %s", &op, &a, &b)
	if err != nil {
		fmt.Println("error: ", err)
		fmt.Println("format: <op> <num1> <num2>")
		return
	}

	if err := validateNum(a); err != nil {
		fmt.Println("first number error:", err)
		return
	}

	if err := validateNum(b); err != nil {
		fmt.Println("second number error:", err)
		return
	}

	fmt.Printf("##%s##%s##%s##", op, a, b)
}
