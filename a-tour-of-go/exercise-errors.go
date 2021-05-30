package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)).
	// https://stackoverflow.com/a/27475316
	// return fmt.Sprintf("cannot Sqrt negative value: %v", float64(e))
	return fmt.Sprintf("cannot Sqrt negative value: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2
	i := 0
	for ; i < 100 && z*z != x; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(x-z*z) <= 0.000000000000001 {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
