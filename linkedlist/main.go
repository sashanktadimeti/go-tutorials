package main
import ("fmt")
type LinkedList struct{
	value int
	next *LinkedList
}
func (ll *LinkedList) PrintList(){
	if ll == nil {
		return
	}
	fmt.Println(ll.value)
	ll.next.PrintList()
}
func main() {
	fmt.Println("This is a placeholder main function.")
}

// Nothing in Go prevents calling a method with nil receiver.