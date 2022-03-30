package main

import "fmt"

// Используeм struct{}, чтобы занимать как можно меньше места
type Set map[string]struct{}

func NewSet(elems []string) Set {
	set := make(Set, len(elems))
	for _, val := range elems {
		set[val] = struct{}{}
	}
	return set
}

func (set *Set) Slice() []string {
	result := make([]string, 0, len(*set))
	for val := range *set {
		result = append(result, val)
	}
	return result
}

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}

	set := NewSet(input)

	fmt.Println(set.Slice())
}
