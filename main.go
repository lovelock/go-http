package main

import (
	"net/http"
	"fmt"
)

type foo struct{}

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/boy":
		fmt.Fprintln(w, "I love you!!!")
	case "/girl":
		fmt.Fprintln(w, "hehe")
	default:
		fmt.Fprintln(w, "Men would stop talking and women would shed tears when they see this. 男默女泪啊哈哈")
	}
}

func main() {
	f := foo{}
	http.ListenAndServe(":8080", f)
}
