package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) push(v T) {
	newNode := &List[T]{next: l.next, val: l.val}
	l.val = v
	l.next = newNode
}

func (l *List[T]) pop() (v T) {
	if l.checkEmpty() {
		fmt.Println("Empty List")
	}
	data := l.val
	l.val = l.next.val
	l.next = l.next.next
	return data
}

func (l *List[T]) print() {
	current := l
	for current != nil {
		fmt.Printf("%v -> ", current.val)
		current = current.next
	}
	fmt.Println(nil)
}

func (l *List[T]) checkEmpty() bool {
	return l == nil
}

func main() {
	fruit := &List[string]{val: "Apple"}

	fruit.push("Banana")
	fruit.push("Pineapple")
	fruit.print()
	fruit.pop()
	fruit.print()
}
