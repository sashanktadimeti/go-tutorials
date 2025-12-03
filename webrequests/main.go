package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const url string = "https://leetcode.com/"
	fmt.Println("http request")
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		} else {
			fmt.Println(string(data))
		}
	}
	defer res.Body.Close()
}
