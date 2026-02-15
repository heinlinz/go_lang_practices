package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 && !ok2 {
			return true
		}

		if v1 != v2 || ok1 != ok2 {
			return false
		}
	}
}

func main() {
	fmt.Print("Comparing 1 and 1: ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED (True)")
	} else {
		fmt.Println("FAILED (False)")
	}

	// Test 2: Different structure (1 vs 2) -> Should be FALSE
	fmt.Print("Comparing 1 and 2: ")
	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("FAILED (True) - Check your Walk function!")
	} else {
		fmt.Println("PASSED (False)")
	}
}
