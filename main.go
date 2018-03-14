package main

import (
	"net"
	"log"
	"fmt"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(conn, "Hello from TCP server")

		conn.Close()
	}
}