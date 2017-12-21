package main

import "fmt"

// Declare the struct data types called rect and circle
type rect struct {
	width, height int
}

type circle struct {
	radius float32
}

// Declare 2 methods named area, one specifically for the rect struct and circle struct
// note that both methods have a pointer receiver
func (r *rect) area() int {
	return r.width * r.height
}

// Pointer receiver allows us to modify the orignal struct data type rather than the copy of the struct.
// this method accepts a pointer of a struct data type, modifies the radius, and outputs the new radius and area
func (r *circle) area() float32 {
	r.radius = 20
	fmt.Println("circle radius changed to: ", r.radius)
	return 3.14 * (r.radius * r.radius)
}

func main() {
	// declare our struct data types
	r := rect{width: 10, height: 5}
	c := circle{radius: 10}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("rect area: ", r.area()) // calculate the area of rect

	fmt.Println("circle original radius is: ", c.radius) // output the original value of c's radius
	fmt.Println("circle area: ", c.area())               // output the radius changed value and the area of c
	fmt.Println("circle new radius is: ", c.radius)      // output c's new radius to confirm that the original value was changed (pointer)

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.

}
