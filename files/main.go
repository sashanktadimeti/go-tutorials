package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "hello"
	file, err := os.Create("./sample.txt")
	if err != nil {
		panic(err)
	} else {
		length, err := io.WriteString(file, content)
		if err != nil {
			panic(err)
		} else {
			fmt.Println(length)
		}
	}
	contentread, err := os.ReadFile("./sample.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(contentread))

}
