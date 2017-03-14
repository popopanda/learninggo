package main

import (
	"fmt"
)

func main() {
	a := 3
	b := 2
	c := "dog"
	d := "cat"

	if a < b {
		fmt.Printf("A is less than B, therefore this is printing %v\n", c)
	} else {
		fmt.Printf("A is greater than B, therefore this is printing %v\n", d)
	}
}
