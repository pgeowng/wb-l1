package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// echo "+ -9_223_372_036_854_775_808 9_223_372_036_854_775_807" | go run main.go

func main() {
	var op, a, b string

	_, err := fmt.Scanf("%s %s %s", &op, &a, &b)
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("format: <op> <num1> <num2>")
		os.Exit(1)
	}

	a64, err := strconv.ParseInt(a, 0, 64)
	if err != nil {
		fmt.Println("first number error:", err)
		os.Exit(1)
	}

	b64, err := strconv.ParseInt(b, 0, 64)
	if err != nil {
		fmt.Println("second number error:", err)
		os.Exit(1)
	}

	bigA := big.NewInt(a64)
	bigB := big.NewInt(b64)

	bigResult := big.NewInt(0)

	switch op {
	case "+":
		bigResult = bigResult.Add(bigA, bigB)
	case "-":
		bigResult = bigResult.Sub(bigA, bigB)
	case "*":
		bigResult = bigResult.Mul(bigA, bigB)
	case "/":
		bigResult = bigResult.Div(bigA, bigB)
	default:
		fmt.Println("unknown operation:", op)
		os.Exit(1)
	}

	fmt.Printf("%s %s %s = %s\n", bigA, op, bigB, bigResult)
}
