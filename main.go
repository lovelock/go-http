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

	http.HandleFunc("/boy", boy)
	http.HandleFunc("/girl", girl)
	http.HandleFunc("/", foo)

	http.ListenAndServe(":8080", nil)
}
