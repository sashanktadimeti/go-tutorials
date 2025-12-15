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
func(ll *LinkedList) sum() int{
	if ll == nil {
		return 0
	}
	return ll.value + ll.next.sum()
}
func main() {
	fmt.Println("This is a placeholder main function.")
	node1 := new(LinkedList)
	node1.value = 10
	node2 := new(LinkedList)
	node2.value = 20
	node3 := new (LinkedList)
	node3.value = 30
	node1.next = node2
	node2.next = node3
	node1.PrintList()
	fmt.Println("Sum:", node1.sum())
}

// Nothing in Go prevents calling a method with nil receiver.
// In this example, we define a simple singly linked list structure with methods to print the list and calculate the sum of its values. The PrintList and sum methods handle nil receivers gracefully, allowing for safe traversal of the list even when it ends with a nil pointer.