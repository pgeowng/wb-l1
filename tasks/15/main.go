package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://research.swtch.com/godata
// Тип string представляет из себя struct{ ptr: *byte, len: uint }
/*
var justString string
func someFunc() {
	// После вызова функции будет выделено в памяти 1024 байт под строку.
	// Даже после отсечения в v[:100], остальная часть останется в памяти.
	// Чтобы избежать этого, можно использовать:
	//   + justString = strings.Clone(v[:100])    // одно выделение памяти
	//   + justString = string([]rune(val)[:100]) // два выделения: string -> []rune, []rune -> string
  v := createHugeString(1 << 10)

  // Строки хранятся в UTF-8.
  // Это означает, что один символ может занимать больше чем один байт.
  // Но при использовании v[:100] происходит отсечение по байтам, исходя из ptr *byte.
  // В итоге, можем получить некоректную строку. Лучше использовать []rune.
  //   ("日本語")[1:2] 							// <?>
  //   string([]rune("日本語")[1:2]) // 本
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

func createHugeString(size int) (result []rune) {
	for i := 0; i < size; i++ {
		var r rune
		switch rand.Intn(3) {
		case 0:
			r = rune('A' + rand.Intn(25))
		case 1:
			r = rune('а' + rand.Intn(33))
		case 2:
			r = rune('あ' + rand.Intn(50))
		}
		result = append(result, r)
	}

	return
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[0:100])
}

func main() {
	rand.Seed(time.Now().UnixNano())
	someFunc()
	fmt.Println(justString)
}
