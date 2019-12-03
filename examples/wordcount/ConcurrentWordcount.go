package main

import (
"fmt"
"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := circle{12}
	fmt.Println("Circle area = ", c.area())

	r := rect{6, 7}
	fmt.Println("Rectangle area = ", r.area())
}