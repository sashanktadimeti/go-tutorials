package main

import "fmt"

func main() {
	num := 10
	if num < 10 {
		fmt.Println("num is less than 10")
	} else if num > 10 {
		fmt.Println("num is greter than 10")
	} else {
		fmt.Println("number is 10")
	}
	// if num := 4; num < 10 {
	// 	fmt.Println(num)
	// 	fmt.Println("number is four")
	//block scope num is being shadowed here similar to js }
	fmt.Println(num)
}
