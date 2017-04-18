package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd, err := exec.Command("echo", "-n", "hello dog\nwoof woof\ncat nip\nmeow meow\nshiba goes woof meow").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", strings.Split(string(cmd), " "))
}
