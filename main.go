package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Longstring string = "hello world!"

func main() {
	var a int = 32
	text := "sashank"
	const b int = 35
	fmt.Printf("the type of b is:%T\n", b)
	fmt.Printf("the type of text is:%T\n", text)
	fmt.Printf("the type of a is: %T\n", a)
	fmt.Println(a)
	fmt.Println(Longstring)
	fmt.Println("performing reading operation")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the number plus one is:", number+1)
	}
	ptr1 := &a
	*ptr1 = *ptr1 * 3
	fmt.Println("the value in the pointer is ", a)
	fmt.Printf("type of ptr1 is %T\n", ptr1)
}
