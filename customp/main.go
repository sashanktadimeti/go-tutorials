package main

import (
	"customp/mathutils"
	"fmt"
)

func main() {
	fmt.Println("helllo...")
	sum := mathutils.Add(5, 10)
	factorial := mathutils.Factorial(5)
	fmt.Println("Sum:", sum)
	fmt.Println("5! (Factorial):", factorial)
}
