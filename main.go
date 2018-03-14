package main

import (
	"net/http"
	"fmt"
)

func boy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I love you!!!")

}
func girl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hehe")

}
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Men would stop talking and women would shed tears when they see this. 男默女泪啊哈哈")
}

func main() {

	http.Handle("/boy", http.HandlerFunc(boy))
	http.Handle("/girl", http.HandlerFunc(girl))
	http.Handle("/", http.HandlerFunc(foo))

	http.ListenAndServe(":8080", nil)
}
