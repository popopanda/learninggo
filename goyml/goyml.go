package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	type Marvel struct {
		Name  string
		Team  string
		Power string
	}

	type Config struct {
		Marvel []Marvel
	}

	var config Config
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

	fmt.Printf("Value: %#v\n", config.Marvel)

}
