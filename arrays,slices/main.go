package main

import (
	"fmt"
	"sort"
)

func main() {
	// var fruitlist = []string{"apple", "banaana", "kaya"}
	// fruitlist=append(fruitlist,"potato","carrot")
	vegetableList := make([]int, 4)
	vegetableList = append(vegetableList, 10, 2, 3, 4, 5, 14, 18)
	sort.Ints(vegetableList)
	fmt.Println(vegetableList)
	vegetableList = append(vegetableList[:2], vegetableList[3:]...)//to remove an element located at index
	fmt.Println(vegetableList)
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	for _, item := range items {
		a = append(a, item[:])   // convert array → slice
	}

	fmt.Println(items)
	fmt.Println(a)

	// fmt.Println(fruitlist)
	// ptr := new(int)
	// fmt.Println(*ptr)
	// ptr1 := make([]int, 4)
	// fmt.Println(ptr1)

}
