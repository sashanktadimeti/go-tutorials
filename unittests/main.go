package main
import ("os"
	"fmt"
	"unittests/hello"
)
func main() {
	if len(os.Args) > 1{
		fmt.Println(hello.SayHello(os.Args[1]))
	}else {
		fmt.Println("hello world")
	}
}