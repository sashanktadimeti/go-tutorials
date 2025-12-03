package main

import "fmt"

func main() {

	for number := 0; number < 10; number++ {
		fmt.Println(number)
	}
	var vegetableList = []int{1, 2, 3, 4, 5, 6, 7}
	for number := 0; number < len(vegetableList); number++ {
		fmt.Println(vegetableList[number])
	}
	for index, number := range vegetableList {
		fmt.Println(index, number)
	}
	roguevalue := 1
	for roguevalue < 10 {
		fmt.Println(roguevalue)
		roguevalue++
	} //simlar to while

}
