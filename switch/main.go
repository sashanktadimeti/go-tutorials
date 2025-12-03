package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := rand.Intn(6)
	number = number + 1
	switch number {
	case 1:
		fmt.Println("i am one")
	case 2:
		fmt.Println("i am two")
	case 3:
		fmt.Println("i am three")
	case 4:
		fmt.Println("i am four")
	case 5:
		fmt.Println("i am five")
	case 6:
		fmt.Println("i am six")
	}
}
