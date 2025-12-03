package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Before sending 5")
		ch <- 5 
		time.Sleep(2 * time.Second)
		fmt.Println("After sending 5")
	}()

	val := <-ch
	fmt.Println("Received:", val)

}
