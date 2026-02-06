package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, y float64
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.y = v.y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.y = v.y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.y*v.y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.y*v.y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("Before scaling: %+v , Abs: %+v\n", v, v.Abs())
	v.scale(12)
	fmt.Printf("After scaling: %+v , Abs: %+v\n", v, v.Abs())
	ScaleFunc(&v, 2)

	p := &Vertex{5, 5}
	p.scale(10)
	ScaleFunc(p, 3)

	// fmt.Println(v, p)
	// fmt.Println(v.Abs())
	// fmt.Println(AbsFunc(v))
	// fmt.Println(p.Abs())
	// fmt.Println(AbsFunc(*p))
}
