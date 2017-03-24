package main

import "fmt"

func main() {
	// firstMap := make(map[string]int)
	// firstMap["shiba"] = 1
	// firstMap["husky"] = 3

	secondMap := map[string]int{
		"shiba":   1,
		"corgi":   2,
		"husky":   3,
		"golden":  4,
		"samoyed": 5,
	}
	fmt.Printf("my secondMap contains: %v\n", secondMap)

	//outputs in random order
	for key, value := range secondMap {
		fmt.Printf("Key: %v, and value is: %v\n", key, value)
	}
	//change value in a map
	secondMap["shiba"] = 10
	fmt.Printf("Shiba now changed to: %v\n", secondMap["shiba"])

	//adding value to a map
	secondMap["cat"] = 500
	fmt.Printf("\nThe new list contains: %v", secondMap)

	//delete a key
	delete(secondMap, "golden")
	fmt.Printf("\nDeleted \"golden\", the map now contains: %v", secondMap)

	//maps are reference types
	//specify size for large maps improve performance
	//make(map[type]<valuetype>, size )
}
