package mathutils

func Factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * Factorial(num-1)
}
