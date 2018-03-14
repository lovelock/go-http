package main

import (
	"net"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	eee, err := ioutil.ReadAll(conn)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(eee))
}
