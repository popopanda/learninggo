package main

import (
	"fmt"
	"strings"
)

func main() {
	myvar1 := "Hello"

	fmt.Println(allCaps(myvar1))
	allTitle("shiba", "corgi", "husky", "golden retriever", "penguins")
}

func allCaps(x string) (y string) {
	y = strings.ToUpper(x)
	return
}

//variadic function
func allTitle(a ...string) string {
	// return strings.Join(a, ",")
	for _, myVar := range a {
		fmt.Println(strings.Title(myVar))
	}
	return ""
}
