package main

import (
	"fmt"
)

func fibonacci() func() int {
	f0, f1 := 0, 1
	return func() int {
		n := f0
		f0, f1 = f1, f0+f1
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
