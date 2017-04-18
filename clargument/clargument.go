package main

import (
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	for index := 1; index < len(os.Args); index++ {
// 		fmt.Println(os.Args[index])
// 	}
// }

// func main() {
// 	s, sep := "", ""
// 	for _, value := range os.Args[1:] {
// 		s += sep + value
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

func main() {

	if len(os.Args) < 2 {
		fmt.Println("There was no arguments passed")
		return
	}

	joinedString := strings.Join(os.Args[1:], " ")
	fmt.Println(joinedString) //join the values in the array seperated by space

	fmt.Println(os.Args[1:]) //array
	fmt.Println("The number of items is: ", len(os.Args[1:]))
}
