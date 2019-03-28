package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalln("there was an issue listening on that port")
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("There was an issue accepting incoming connections")
		}

		go serve(conn)

		fmt.Println("Code got here")
	}
}

func serve(connection net.Conn) {

	scanner := bufio.NewScanner(connection)

	i := 0

	var method string
	var uri string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			fields := strings.Fields(ln)

			method = fields[0]
			uri = fields[1]
		}

		if ln == "" {
			fmt.Println("You've reached the end of the scan!")
			break
		}

		i++
	}

	payload := `<h2>Home page</h1>
	<form>
		<input ></input>
	</form>
	<a href="apply.gohtml">/apply</a>`

	io.WriteString(connection, "HTTP/1.1 200 OK\r\n")

	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(payload))

	fmt.Fprint(connection, "Content-Type:text/html\r\n")

	io.WriteString(connection, "\r\n")

	io.WriteString(connection, payload)

	fmt.Fprintf(connection, "<h2>Method: %v</h2>", method)

	fmt.Fprintf(connection, "<h2>URI: %v</h2>", uri)

	connection.Close()

}
