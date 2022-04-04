package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http

func main() {
	resp, err := http.Get("https://www.liwenzhou.com")
	if err != nil {
		fmt.Println("get url failed, err:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get body failed, err:", err)
		return
	}
	fmt.Println(string(b))
}
