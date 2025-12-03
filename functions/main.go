package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	result := adder(2, 3, 5)
	result1 := adderwithmanyargs(3, 4, 5, 6, 7, 8)
	fmt.Println(result, result1)
	result2, message := multreturnvalues(10, 11)
	fmt.Println(result2, message)
}
func adder(num1 int, num2 int, num3 int) int {
	return num1 + num2 + num3
}
func adderwithmanyargs(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
} //pro functions
func multreturnvalues(val1 int, val2 int) (int, string) {
	return val1 + val2, "resukt calculated"
}
