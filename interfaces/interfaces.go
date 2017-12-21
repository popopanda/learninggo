package main

import "fmt"

type Calculations interface {
	area() float64
}

type Shape struct {
	Sides int
}

type Circle struct {
	Shape
	Radius float64
}

type Square struct {
	Shape
	Length float64
	Width  float64
}

func (s Square) area() float64 {
	return s.Length * s.Width
}

func (c Circle) area() float64 {
	return 3.14 * (c.Radius * c.Radius)
}

func calculate(a Calculations) {
	fmt.Println(a.area())
}

func main() {
	circ := Circle{Shape{0}, 82.12}
	squa := Square{Shape{4}, 12.1, 12.1}

	calculate(circ)
	calculate(squa)
}
