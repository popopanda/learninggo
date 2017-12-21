package main

import (
	"fmt"
)

// defining a "base" animal struct
type animal struct {
	name     string
	eatMeat  bool
	isMammal bool
}

// type dog and cat struct receives values from animal struct (think of inherenting)
type dog struct {
	animal
	color string
}

type cat struct {
	animal
	color string
}

// defining speak method for cat and dog struct
func (l *cat) Speak() {
	fmt.Printf("%v says: meow\n", l.name)
}

func (l *dog) Speak() {
	fmt.Printf("%v says: woof\n", l.name)
}

func main() {
	kobe := new(dog)
	kobe.name = "Sir Kobe"
	kobe.eatMeat = true
	kobe.isMammal = true
	kobe.color = "white"

	lambchops := new(cat)
	lambchops.name = "Sir Lambchops"
	lambchops.eatMeat = true
	lambchops.isMammal = true
	lambchops.color = "black"

	//confirming structs were initialized properly
	// fmt.Println(lambchops.name, lambchops.color, lambchops.eatMeat)

	// invoking the speak method corresponding to the correct data type for the methods
	kobe.Speak()
	lambchops.Speak()
}
