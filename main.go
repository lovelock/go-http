package main

import (
	"net/http"
	"fmt"
)

type (
	boy struct{}
	girl struct {
	}
	foo struct {
	}
)

func (b boy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I love you!!!")

}
func (g girl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hehe")

}
func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Men would stop talking and women would shed tears when they see this. 男默女泪啊哈哈")
}

func main() {

	b, g, f := boy{}, girl{}, foo{}

	mux := http.NewServeMux()
	mux.Handle("/boy", b)
	mux.Handle("/girl", g)
	mux.Handle("/", f)

	http.ListenAndServe(":8080", mux)
}
