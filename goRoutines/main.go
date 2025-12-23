package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Before sending 5")
		ch <- 5 
		time.Sleep(2 * time.Second)
		fmt.Println("After sending 5")
	}()
	val := <-ch
	wg.Wait()
	fmt.Println("Received:", val)

}
