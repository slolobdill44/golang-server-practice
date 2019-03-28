package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
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

	var payload string

	switch uri {
	case "/":
		payload = http.HandlerFunc(home)
	case "/apply":
		payload = `
		<h2>Apply Page</h2>
		<form>
			<input ></input>  <button>submit</submit>
		</form>

		<p><a href="home">/home</a></p>`
	default:
		payload = `
		
		<h2>Home page</h1>
		<a href="apply">/apply</a>`
	}

	io.WriteString(connection, "HTTP/1.1 200 OK\r\n")

	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(payload))

	fmt.Fprint(connection, "Content-Type:text/html\r\n")

	io.WriteString(connection, "\r\n")

	io.WriteString(connection, payload)

	fmt.Fprintf(connection, "<h2>Method: %v</h2>", method)

	fmt.Fprintf(connection, "<h2>URI: %v</h2>", uri)

	connection.Close()

}

func home(resp http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("home.gohtml")

	if err != nil {
		log.Fatalln("There was an error parsing the home page")
	}

	err = tpl.ExecuteTemplate(resp, "home.gohtml", nil)
}
