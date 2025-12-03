package main
import ("fmt")
func main(){
	a:= []func(){}
	for i:=0;i<5;i++{
		a = append(a, func(){
			fmt.Println(i)
		})
	}
	for _,f := range a{
		f()
	}
}