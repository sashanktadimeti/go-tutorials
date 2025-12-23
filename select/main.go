package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{make(chan int), make(chan int)}
	for i := range chans {
		go func(i int, ch chan<- int){
			defer close(ch)
			for {
			if i==1 {
				time.Sleep(time.Duration(i+23) * time.Second)
			}
			ch<-(i+1)
			}
		}(i,chans[i])
	}
	for {select{
		case m1,ok:=<-chans[0]:
			if !ok{
				return
			}
			log.Println("from channel 1",m1)
		case m2,ok:=<-chans[1]:
			if !ok{
				return
			}
			log.Println("from channel 2",m2)
	}
	}
}

	// for{
	// m1 := <-chans[0]
	// log.Println("from channel 1",m1)
	// m2 := <- chans[1]
	// log.Println("from channel 2",m2)
	// }