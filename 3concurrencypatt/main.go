package main

import (
	"fmt"
	"time"
	"sync"
)

type item struct {
	price    int
	category string
}

func main() {
	c := gen(
		item{10, "A"},
		item{20, "B"},
		item{30, "C"},
		item{40, "D"},
	)

	out := discount(c)

	for it := range out {
		fmt.Println(it)
	}
}

func discount(items <-chan item) <-chan item {
	out := make(chan item)

	go func() {
		defer close(out)
		time.Sleep(time.Second/2)
		for it := range items {
			if it.category == "A" {
				it.price /= 2
			}
			out <- it
		}
	}()

	return out
}

func gen(items ...item) <-chan item {
	out := make(chan item, len(items))

	for _, it := range items {
		out <- it
	}
	close(out)
	return out
}
 func fanIn(channels ... <-chan item) <- chan item{
	var wg sync.WaitGroup
	out := make(chan item)
	wg.Add(len(channels))
	output := func(c <- chan item){
		defer wg.Done()	
		for it := range c {
			out <- it
		}			
	}
	for _,c := range channels {
		go output (c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
 }
