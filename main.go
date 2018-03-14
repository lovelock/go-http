package main

import (
	"net/http"
	"fmt"
)

type foo struct{}

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Implement the Handler interface")
}

func main() {
	f := foo{}
	http.ListenAndServe(":8080", f)
}
