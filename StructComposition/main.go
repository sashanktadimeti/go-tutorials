package main

import ("fmt"
	"path/filepath"
)

type Pair struct {
	length string
	width  string
}
type Pairheight struct{
	Pair
	height string
}
func(p Pairheight) Area() string {
	return fmt.Sprintf("Given length %s, breadth %s and height %s",p.length,p.width,p.height)
}
func (p Pair) Area() string {
	return fmt.Sprintf("Given length %s and breadth %s",p.length,p.width)
}
func Filename(p Pair) string {
	return filepath.Base(p.length)
}
type Filenamer interface{
	Filename() string
}
func (p Pair) Filename() string {
	return filepath.Base(p.length)
}
func main() {
	p := Pairheight{Pair{"10cm","5cm"},"15cm"}
	fmt.Println(p.Area())
	// fmt.Println(Filename(p)) this throws an error because Pairheight does not implement Pair
	fmt.Println(Filename(p.Pair)) // this works
	var fn Filenamer = Pairheight{Pair{"69cm","10cm"},"25cm"}
	fmt.Println(fn.Filename())

}
// In this example, we demonstrate struct composition in Go. The Pairheight struct composes the Pair struct, inheriting its fields and methods. We define an Area method for both Pair and Pairheight structs. When we call the Area method on a Pairheight instance, it uses the fields from both the embedded Pair struct and its own height field to provide a more detailed output.
// This showcases how struct composition allows for code reuse and extension of functionality in Go.