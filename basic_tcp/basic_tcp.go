package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleScan(conn)
	}
}

func handleScan(conn net.Conn) {
	defer conn.Close()

	receive(conn)

}

func receive(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			m := strings.Fields(ln)[0]
			p := strings.Fields(ln)[1]
			fmt.Println("**METHOD", m)
			fmt.Println("**PATH", p)
			respond(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn, ln string) {

	method := strings.Fields(ln)[0]
	uriPath := strings.Fields(ln)[1]

	if uriPath == "/" && method == "GET" {
		index(conn)
	}
	if uriPath == "/about" && method == "GET" {
		about(conn)
	}

}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Sup World</strong><p><a href="/about">About</a></p></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>What about this server? It's a great learning tool</strong><p><a href="/">Go home</a></p></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
