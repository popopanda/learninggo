package main

import (
	"fmt"
	"strings"
)

func aMsg() func(y string) string {
	return func(y string) string {
		adjY := strings.ToUpper(y)
		return adjY
	}
}

func main() {

	func(m string) {
		fmt.Println(m)
	}("world")

	textone := aMsg()
	fmt.Println(textone("hello"))
}
