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

//concurrency is the independently executing processes

//goroutines are "threads" that lay on top of system threads. more lightweight and avoid overhead of thread switches
//Main function is like a goroutine

//Channels
// unbuffered channel doesnt specify a size (ie myChannel := make(chan int)), this causes the goroutine that is sending data to lock until there is another availabe goroutine to retreive the data
// buffered channel specify the size (ie. myChannel := make(chan int, 5)).
//5 representing how many data items the channel and hold.Doesnt lock up goroutine sending data.
//Sending go routine can put data onto the channel and continue with its task.
//As long as channel is not full, goroutines can continue dropping data onto the channel.
//If buffered channel is full, sending goroutine will be locked until the channel is full
// buffered or unbuffered, if there is no data in the channel, and a receiving goroutine has no data to grab, it will lock.
