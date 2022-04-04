package main

import (
	"fmt"
	"net/http"
)

/// net/http server

func sayHello(resp http.ResponseWriter, r *http.Request) {
	fmt.Fprint(resp, "舔得好!")
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe("192.168.16.43:8080", nil)
	if err != nil {
		fmt.Println("start http server failed,err:", err)
		return
	}
}
