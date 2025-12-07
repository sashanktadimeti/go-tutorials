package main
import (
	"strconv"
	"strings"
	"fmt"
)
type IntSlice [] int
func (is IntSlice) String() string {
	var strs []string
	for _,v := range is{
		strs = append(strs, strconv.Itoa(v))
	}
	return "[[" + strings.Join(strs, ", ") + "]]"
}
func main() {
	s := IntSlice{1,2,3,4,5}
	// fmt.Println(s.String())
	fmt.Printf("%T %[1]v\n", s)
}


//  methods can be attached to any named type, including types defined based on built-in types like int, string, or slices. In this example, we define a named type IntSlice based on []int and attach a String method to it. This method converts the slice of integers into a formatted string representation.