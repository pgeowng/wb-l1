package main

import (
	"fmt"
	"reflect"
)

//　Переменная типа interface{...} хранит в себе:
//    + Значение присвоенной переменной.
//    + Тип присвоенной переменной.
// Для получение типа специально создан reflect.
// https://research.swtch.com/interfaces
// https://go.dev/blog/laws-of-reflection

type Typer interface {
	Type() string
}

type Node struct{}

func (n Node) Type() string {
	return "Node"
}

// Изпользуем reflect, чтобы рассмотреть полученный тип в interface{}
// Данное решение является более медленным, чем type-assertion/type-switch
// из-за дополнительных выделений памяти.
func GetType(any interface{}) (result string) {
	defer func() {
		if err := recover(); err != nil {
			result += " err: " + reflect.ValueOf(any).String()
		}
	}()

	typeof := reflect.TypeOf(any)
	typeofStr := "nil"
	kindofStr := "nil"
	if typeof != nil {
		typeofStr = typeof.String()
		kindofStr = typeof.Kind().String()
	}

	return fmt.Sprintf("%11s %6s", typeofStr, kindofStr)
}

func main() {
	var n Node
	var pn *Node
	var ss Typer
	var ssn Typer = &n
	var i int
	var s string
	var b bool
	var si []int
	var ci chan int
	var mii map[int]int

	fmt.Printf("%-4s %-11s %-6s\n", "var", "type", "kind")
	fmt.Println("n:  ", GetType(n))
	fmt.Println("pn: ", GetType(pn))
	fmt.Println("ss: ", GetType(ss))
	fmt.Println("ssn:", GetType(ssn))
	fmt.Println("i:  ", GetType(i))
	fmt.Println("s:  ", GetType(s))
	fmt.Println("b:  ", GetType(b))
	fmt.Println("si: ", GetType(si))
	fmt.Println("ci: ", GetType(ci))
	fmt.Println("mii:", GetType(mii))
}
