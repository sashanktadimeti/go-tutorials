package main

import (
	"fmt"
	"sort"
)

func main() {
	// var mymap map[string]int - it is nil map, cannot add key-value pairs, doesnt point to any memory location
	// mymap := map[string]string{} - it is initialized map, can add key-value pairs
	mymap := make(map[string]string)
	// value,ok := mymap["JS"](returns zero value and false if key not present)
	// if ok {
	// 	fmt.Println("Key is present with value:", value)
	// } else {
	// 	fmt.Println("Key not present")
	// }
	mymap["JS"] = "javascript"
	mymap["PY"] = "Python"
	mymap["RB"] = "Ruby"
	// for key, value := range mymap {
	// 	fmt.Printf("For key is %v\n ,value is %v\n", key, value)
	// }
	// for key, value := range mymap {
	// 	fmt.Printf("the key is %v and value is %v\n", key, value)
	// }
	newmap := map[string]int{}
	newarr := []string{"apple", "banana", "grapes", "apple", "banana", "orange"}
	for i:=0; i<len(newarr);i++ {
		newmap[newarr[i]]++
	}
	type newstruct struct {
		Name  string
		Count int
	}
	var newstructs []newstruct
	for key,value := range newmap{
		newstructs = append(newstructs, newstruct{key,value})
	}
	sort.Slice(newstructs, func(i,j int) bool {
		return newstructs[i].Count > newstructs[j].Count
	})
	for value := range(newstructs){
		fmt.Printf("the key is %v and value is %v\n", newstructs[value].Name, newstructs[value].Count)
	}
}
