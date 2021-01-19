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

// The square and circle struct must satisfy the method of the calculations interface, which in this case is only area()
// Meaning, it must also have an area() method in order to be part of the calculations interface. 
// If the calculation interface also defines other methods, the data types must also define it to be part of the interface.

func (s Square) area() float64 {
	return s.Length * s.Width
}

func (c Circle) area() float64 {
	return 3.14 * (c.Radius * c.Radius)
}

// additional methods can be defined for the data type that arent part of the interface.
// For example, the circle struct can also define a circumference method, which isn't specified by the calculations interface.  
// But it already satisfies the calculations method because it has the area() method.

// See type assertion to call the circumference of the area method (converting the calculation interface type back to the circle data type to call the circumference method)

func calculate(a Calculations) {
	fmt.Println(a.area())
}

func main() {
	circ := Circle{Shape{0}, 82.12}
	squa := Square{Shape{4}, 12.1, 12.1}

	calculate(circ)
	calculate(squa)
}
