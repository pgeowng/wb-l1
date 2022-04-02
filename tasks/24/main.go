package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewZeroPoint() Point {
	return Point{0, 0}
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) Distance(o Point) float64 {
	dx := o.x - p.x
	dy := o.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p Point) Shift(x, y float64) Point {
	return Point{p.x + x, p.y + y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.x, p.y)
}

func main() {
	o := NewZeroPoint()
	a := NewPoint(3, 4)

	fmt.Println(o, a, o.Distance(a), o.Shift(8, 6).Distance(o))
}
