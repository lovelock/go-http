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

		handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()

	respond(conn)
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html><head><meta charset="UTF-8"><title>TCP Example</title></head><body>Hello from TCP server</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}