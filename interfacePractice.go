package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height, width float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var s Shape
	r := Rectangle{4.5, 5.5}
	c := Circle{3.5}

	s = r
	describe(s)
	fmt.Printf("The area of the circle is %v\n", s.Area())

	s = c
	describe(s)
	fmt.Printf("The area of the circle is %v", s.Area())
}

func describe(s Shape) {
	fmt.Printf("(%v, %T\n", s, s)
}
