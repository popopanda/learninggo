package main

import "fmt"

// Returns another function, which is defined anonymously.
// The returned function closes over the variable i to form a closure
func intSeq() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	increase := intSeq()

	fmt.Println(increase())
	fmt.Println(increase())
	fmt.Println(increase())
	fmt.Println(increase())
	fmt.Println(increase())

	newIncrease := intSeq()
	fmt.Println(newIncrease())
}
