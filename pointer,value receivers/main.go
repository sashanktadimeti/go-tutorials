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
}