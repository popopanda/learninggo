package main

import "fmt"

func main() {
	myDogs := []string{"shiba inu", "corgi", "husky", "samoyed"}
	myDreamDogs := []string{"shiba inu", "chow chow", "corgi", "golden retriever", "akita"}

	for _, i := range myDogs {
		for _, b := range myDreamDogs {
			if i == b {
				fmt.Printf("%v from myDogs matches with %v in myDreamDogs\n", i, b)
			}
		}
	}
}
