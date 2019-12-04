package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Type() string
	CalculateArea() float64
	CalculatePerimeter() float64
}

type Circle struct {
	radius float64
	Shape
}

func (c Circle) Type() string {
	return "Circle"
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) CalculatePerimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectanle struct {
	height, width float64
	Shape
}

func (r Rectanle) Type() string {
	return "Rectangle"
}

func (r Rectanle) CalculateArea() float64 {
	return r.width * r.height
}

func (r Rectanle) CalculatePerimeter() float64 {
	return (r.height + r.width) * 2
}

func PrintMeasures(s Shape) {
	fmt.Println("Shape: ", s.Type())
	fmt.Println("----------------------------")
	fmt.Println("    Area: ", s.CalculateArea())
	fmt.Println("    Perimeter: ", s.CalculatePerimeter())
	fmt.Println("----------------------------")
}

func main() {
	c := Circle{
		radius: 12,
		Shape:  nil,
	}
}

