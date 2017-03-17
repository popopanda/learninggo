package main

import (
	"fmt"
)

func main() {
	// myDogs := make([]string, 2, 4)
	myDogs := []string{"corgi", "husky", "shiba inu", "samoyed"}

	fmt.Printf("Length is: %d. \nCapacity is %d\n", len(myDogs), cap(myDogs))

	fmt.Println(myDogs[:2])

	myDogs[0] = "shiba inu"

	fmt.Println(myDogs)

	myNewDogs := myDogs[:2]

	fmt.Println(myNewDogs)

	myNewDogs[0] = "cat"

	fmt.Println(myDogs)
	fmt.Println(myNewDogs)
}
