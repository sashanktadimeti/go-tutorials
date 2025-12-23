package main

import (
	"fmt"
	"time"
)

const (
	concurrencyLevel = 3
)

var (
	requestIds = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

func main() {
	queue := make(chan bool, concurrencyLevel)
	// var wg sync.WaitGroup
	for _, id := range requestIds {
		queue <- true
		fmt.Println("sashankid:", id)
		// wg.Add(1)
		go func(id int) {
			defer func() {
				<-queue
			}()
			// defer wg.Done()
			makerequest(id)
		}(id)
	}
	// wg.Wait()
	// for i:=0;i<concurrencyLevel;i++ { 
	// 	queue<-true
	// }  fix with a for loop
}
func makerequest(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println(id)
}