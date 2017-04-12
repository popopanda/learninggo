package main

import "fmt"

var a = 1
var b = 2
var c = "ten"
var cPTR = &c

func main() {
	// a := 10
	// b := "golang"
	// c := true
	// d := "adsfasdf"
	// firstname, lastname := "Jeffrey", "leung"

	// fmt.Printf("%v \n", a)
	// fmt.Printf("%v \n", b)
	// fmt.Printf("%v \n", c)
	// fmt.Println(d)

	// fmt.Printf("My name is %v %v", firstname, lastname)
	b := "20"
	fmt.Println(a)
	fmt.Println(&b)

	fmt.Println("pointer reference is:", cPTR, "and value of cPTR is:", *cPTR)

	z := 10
	x := z
	z = 50

	fmt.Printf("value of z: %v, value of x: %v\n", &z, &x)

	fmt.Printf("z is changed to %v and x is %v\n", z, x)

}
