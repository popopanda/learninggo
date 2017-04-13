package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	type Marvel struct {
		Marvel map[string][]string
	}

	var config Marvel
	ymlFile, err := ioutil.ReadFile("./marvel.yml")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(ymlFile, &config)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for _, hero := range config.Marvel {
		for _, i := range hero {
			fmt.Println(i)
		}
		fmt.Println()
		// fmt.Println(strings.Join(hero, ", "))
		// fmt.Println(hero)
	}

	// fmt.Printf("Value: %#v\n", config.Marvel)

}
