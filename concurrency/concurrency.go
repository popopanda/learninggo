package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("hello world")
	}()

	go func() {
		defer waitGroup.Done()
		fmt.Println("goodbye world")
	}()

	waitGroup.Wait()
}

//results - unsure why its different:

// $ go run concurrency.go
// 1
// goodbye world
// 2
// 3
// 4
// 5
// hello world
// $ go run concurrency.go
// goodbye world
// 1
// 2
// 3
// 4
// 5
// hello world
