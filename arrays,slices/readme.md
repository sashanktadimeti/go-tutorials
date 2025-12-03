Keeping a pointer to an element of slice is always risky
type user struct{
    count int;
    name string
}
func icremenent(memeber *user){
    user.count++
}
func main(){
    a := []user{{"sashank",0}, {"skanda",1}}
    b := &user[0]
    c := user{"chiru",2}
    a := append(a,c)
    increment(b)
    //sashank's count will be zero because reallocation of a to new mwmory address can happend and b no longer points to the actual location (kind of a dangling pointer)
}