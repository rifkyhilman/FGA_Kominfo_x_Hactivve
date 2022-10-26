package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangel struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangel) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangel) perimeter() float64 {
	return 2 * (r.height * r.width)
}

func main() {
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangel{width: 3, height: 2}

	fmt.Printf("Type of c1: %T", c1)
	fmt.Printf("Type of c1: %T", r1)
}
