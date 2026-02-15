package main

import (
	"fmt"
)

type ErrNegativeV float64

func (e ErrNegativeV) Error() string {
	return "cannot sqrt a negative value"
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeV(x)
	}

	z := float64(1)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (z * 2)
		fmt.Println(z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
