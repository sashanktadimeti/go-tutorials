package main
import ("fmt"
"net/http"
"time"
)
type result struct{
	url string
	err error 
	latency time.Duration
}
func get(url string, ch chan<- result){
	start := time.Now()
	// perform http get request
	if resp,err := http.Get(url); err!=nil{
		ch <- result{url,err,0}
	}else{
		defer resp.Body.Close()
		t := time.Since(start)
		ch <-result{url,nil,t}
	}
}
func main(){
	urls := []string{"http://google.com","http://facebook.com","http://twitter.com"}
	channel := make(chan result)
	for _,url := range urls{
		go get(url, channel)
	}
	for i := 0; i < len(urls); i++ {
		r := <-channel
		fmt.Printf("URL: %s, Error: %v, Latency: %v\n", r.url, r.err, r.latency)
	}

}