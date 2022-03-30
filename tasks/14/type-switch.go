package main

import (
	"fmt"
)

type Typer interface {
	Type() string
}

type Node struct{}

// Без указателя при проверки *Node,
// Go попытается разыменовать указатель и вызвать метод func(n Node) Type.
// Поэтому стандартными средствами нельзя определить указатель в переменной interface{}
func (n *Node) Type() string {
	return "*Node"
}

// Используем type-switch для определения типа.
// Данный способ не может определить вложенный тип slice, map, chan, *T,
// если это не задано явно.
func GetType(any interface{}) string {
	if val, ok := any.(Typer); ok {
		return "Typer(" + val.Type() + ")"
	}

	switch v := any.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int(" + fmt.Sprint(len(v)) + ")"
	case []interface{}:
		return "slice of interface{}"
	case map[interface{}]interface{}:
		return "map[interface{}]interface{}"
	default:
		return "unknown type"
	}
}

func main() {
	var n Node
	var pn *Node
	var ss Typer
	var i int
	var s string
	var b bool
	var si []int
	var ci chan int
	var mii map[int]int

	fmt.Println("n:  ", GetType(n))
	fmt.Println("pn: ", GetType(pn))
	fmt.Println("ss: ", GetType(ss))
	fmt.Println("i:  ", GetType(i))
	fmt.Println("s:  ", GetType(s))
	fmt.Println("b:  ", GetType(b))
	fmt.Println("si: ", GetType(si))
	fmt.Println("ci: ", GetType(ci))
	fmt.Println("mii:", GetType(mii))
}
