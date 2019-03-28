package main

import (
	"io"
	"net/http"
)

func dog(resp http.ResponseWriter, req *http.Request) {

	io.WriteString(resp, "woof woof woof")

}

func cat(resp http.ResponseWriter, req *http.Request) {

	io.WriteString(resp, "meow meow meow")
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8081", nil)
}
