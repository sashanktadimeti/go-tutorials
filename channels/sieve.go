package main
import ("fmt")
func generate(ch chan<- int){
	for i:=2; ;i++{
		ch<-i
	}
	close(ch)
}
func filter(ch <-chan int, out chan<- int , prime int){
	for i:= range ch{
		if i%prime!=0 {
			out<-i
		}
	}
	close(out)
}
func main(){
	fmt.Println("finding primes using sieve")
	ch := make(chan int)
	go generate(ch)
	for i:=1;i<10;i++ {
		prime,ok:=<-ch
		if !ok {
			break
		}
		fmt.Println(prime);
		out:=make(chan int)
		go filter(ch,out,prime)
		ch=out
	}
}