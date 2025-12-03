var s[] int - nil slice // underlying array is not allocated len,cap - 0,0,true
t : = []int{}  // underlying array is allocted len,cap - 0,0,false , {}
u := make([]int,5) // [0,0,0,0,0] len,cap - 5,5,false
v := make([]int,0,5) // [] len,cap - 0,5,false,{}
// appending to a nil slice is fyn unlike nil map

WTF moments:
        func main(){
            a := [3]int{1,2,3}
            b := a[:1] // len(b)-1 , cap(b)-3
            c := b[:2] // len(c)-2 , cap(c)-3
            // the value of c is {1,2}
        }
        from go1.2
        d := a[0:1:1] // [i:j:k](len-(j-i), cap - (k-i)) // three index operator
        d has only [1]