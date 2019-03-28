package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello I am your TCP server")
		io.WriteString(conn, "\nI am here to serve you a website")
		io.WriteString(conn, "\nI will provide you with HTML soon\n")

		conn.Close()
	}
}
