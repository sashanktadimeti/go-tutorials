package main

import "fmt"
// defer copies values not the references(except for closures)-it is function scoped
func main() {
	fmt.Println("defer in go lang")
	fmt.Println("helllo")
	defer mydefer()
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	defer fmt.Println("four")
}
func mydefer() {
	for number := 0; number < 5; number++ {
		defer fmt.Println(number)
	}
}
