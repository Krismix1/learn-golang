package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// try removing the pointer from the type, and notice how the value changes.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Modify(v Vertex) {
	v.X = v.X - 1
	v.Y += v.X
}

func main() {
	v := Vertex{1, 2}
	v.Scale(10)
	fmt.Println(v.Abs())
	Modify(v)
	fmt.Println(v)
}
