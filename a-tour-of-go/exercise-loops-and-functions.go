package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	i := 0
	for ; i < 100 && z*z != x; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(x-z*z) <= 0.000000000000001 {
			break
		}
		// fmt.Println(fmt.Sprintf("x = %v, z = %v", x, z))
	}
	fmt.Println(i)
	return z
}

func main() {
	for i := 1; i < 20; i++ {
		fmt.Printf("math.Sqrt(%v) = %v, Sqrt(%v) = %v\n", i, math.Sqrt(float64(i)), i, Sqrt(float64(i)))
	}
}
