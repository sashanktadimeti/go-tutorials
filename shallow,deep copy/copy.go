package main
import ("fmt")
func main(){
	arr:=[]int{1,2,3,4,5};
	arr1 := make([]int, len(arr))
	copy(arr1,arr)
	arr1[0]=90
	fmt.Println("Original Array:",arr)
	fmt.Println("Copied Array:",arr1)
	a := [][]int{{1, 2}, {3, 4}}
	b := make([][]int, len(a))
	// copy(b, a)

	// b[1][1] = 99
	// fmt.Println(a,b) // [[99 2] [3 4]] ❗ inner slice still shared
	for i:= range a {
		b[i] = make([]int,len(a[i]))
		copy(b[i],a[i])
	}

}