package main
import ("fmt")
type point struct{
	x,y float32
}
func (p point) ValueReceiver() {
	fmt.Printf("Value receiver called. x: %f, y: %f\n", p.x, p.y)
}
func (p *point) PointerReceiver() {
	fmt.Printf("Pointer receiver called. x: %f, y: %f\n", p.x, p.y)
}
func main() {
	p := point{10,20}
	p.ValueReceiver()
	p.PointerReceiver() // Go automatically takes the address of p to match the pointer receiver it gets converted to &p

	pp := new(point)
	pp.x = 30
	pp.y = 40
	pp.ValueReceiver() // Go automatically dereferences pp to match the value receiver it gets converted to *pp
	pp.PointerReceiver()
	// don't have to explicitly use & or * when calling methods
	// Go handles the conversion automatically
	// This makes method calls more convenient and less error-prone.
	// you are not supposed to define same method with both pointer and value receivers on the same type
	// it makes go confused about which method to call
	//L-Value , R-value
	// L-value works with pointer receiver but r value does not eg:Point{10,20}.PointerReceiver() will throw an error
	// R-value works with value receiver eg: Point{10,20}.ValueReceiver() works fine
	// a struct implementing a pointer receiver should have all of its methods defined with pointer receivers to avoid confusion and ensure consistent behavior.
}

// method values
// func (p point) Distance(q point) float64 {
// 	return math.Sqrt((p.x - q.x)*(p.x - q.x) + (p.y - q.y)*(p.y - q.y))	
// }
// p:= point{10,20}
// d1 := p.Distance
// q:= point{30,40}
// fmt.Println("Distance:",d1(q)) // calls p.Distance(q)