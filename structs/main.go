package main

import "fmt"

func main() {
	sashank := User{"sashank", 60}
	fmt.Println(sashank)
}

type User struct {
	Name   string
	Rollno int
}
