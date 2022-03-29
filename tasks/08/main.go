package main

import (
	"errors"
	"fmt"
	"strconv"
)

func readNumber() (number int64, err error) {
	fmt.Print("enter number(int64): ")

	var numberStr string
	_, err = fmt.Scanf("%s\n", &numberStr)

	if err != nil {
		return
	}

	number, err = strconv.ParseInt(numberStr, 10, 64)
	return
}

func readIntR(left, right int, msg string) (number int, err error) {
	fmt.Print(msg)
	_, err = fmt.Scanf("%d\n", &number)
	if err != nil {
		return
	}

	if number < left || number > right {
		errMsg := fmt.Sprintf("number is not in range from %d to %d", left, right)
		err = errors.New(errMsg)
	}

	return
}

func main() {
	// Считывание и обработка ошибок
	number, err := readNumber()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	bit, err := readIntR(0, 63, "enter bit(0..63): ")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	value, err := readIntR(0, 1, "enter value(0..1): ")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("before: %064s %d\n", strconv.FormatUint(uint64(number), 2), number)

	// Сначала ставим 0, потом value
	number = number & ^(int64(1)<<bit) | (int64(value) << bit)

	fmt.Printf("after:  %064s %d\n", strconv.FormatUint(uint64(number), 2), number)
}
