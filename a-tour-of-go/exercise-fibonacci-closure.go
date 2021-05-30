package main

import "fmt"

func Fibonacci() func() int {
	x := 0
	y := 1
	z := 0
	return func() int {
		z, x, y = x, y, x+y
		return z
	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
